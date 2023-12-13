//SPDX-License-Identifier: MIT
pragma solidity ^0.8.3;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract NFTCollateralLoan {

    struct LoanProposal {

        // Borrower details
        address borrower;   // THe person floating the loan
        address nftContractAddress; // NFTs supposedly are part of separate contracts.
        uint256 nftTokenId; // What is the ID for the NFT under the contract

        // Loan details 
        uint256 loanAmount;
        uint256 interestRate;
        uint256 duration; // seems like the standard practise is to denote time in seconds 
        uint256 dueDate;  // seconds from epoch ?
        uint256 startTime; // loan start time. time when the lending happens
        
        address lender; // Person fine with giving the loan 

        // Proposal status details
        bool isAccepted; // lender accepted it
        bool isPaidBack; // borrowr paid back
    }

    uint256 public nextProposalId; 

    mapping(uint256 => LoanProposal) public proposals; // a map is better than an array because of the retraceProposal use case. 

    // the examples showed a balance ledger typically used per address .. any need for this ?
    // How useful are events ?

   
    function submitProposal(address _nftContractAddress, uint256 _nftTokenId, uint256 _loanAmount, uint256 _interestRate, uint256 _duration) external {
        IERC721 nftContract = IERC721(_nftContractAddress);
        require(nftContract.ownerOf(_nftTokenId) == msg.sender, "Not the NFT owner");
        nftContract.transferFrom(msg.sender, address(this), _nftTokenId);

        LoanProposal memory newProposal = LoanProposal({
            borrower: msg.sender,
            nftContractAddress: _nftContractAddress,
            nftTokenId: _nftTokenId,
            loanAmount: _loanAmount,
            interestRate: _interestRate,
            duration: _duration,
            dueDate: block.timestamp + _duration,
            startTime: 0,
            lender: address(0),
            isAccepted: false,
            isPaidBack: false
        });


        proposals[nextProposalId++] = newProposal; // any added benefit to using a random id ?
    }
  

    // acceptProposal accepts a proposal for a loan and transfer money  
    function acceptProposal(uint256 _proposalId) external payable {

        LoanProposal storage proposal = proposals[_proposalId]; // needs to be storage for efficienct since we are referring to a stored value in block chain

        // seems like we need this check , when someone accepts the proposal they pay their ether and it should be equal to the amount of loan ?
        require(msg.value == proposal.loanAmount, "Incorrect loan amount");
        
        // if proposal is already accepted, we should give up. This is an important part. 
        // the transaction nature of block chain should ensure that there are no two lenders updating the global state 
        require(proposal.isAccepted == false, "Proposal already accepted");
      
        // Set lender to the person who raised the accept transaction
        proposal.lender = msg.sender;
        proposal.isAccepted = true;
        
         // Transfer
        bool lent = payable(proposal.borrower).send(msg.value);
        require(lent, "Failed to lend");

         proposal.startTime = block.timestamp;

    }

    // calculateInterest calculates the loan interest as an annual %
    function calculateInterest(LoanProposal storage _proposal) private view returns (uint256) {
        // Need to convert everything to seconds
        uint256 timeElapsed = block.timestamp - _proposal.startTime; // Time elapsed in seconds
        uint256 oneYear = 365 * 24 * 60 * 60; // Number of seconds in a year
        uint256 interest = (_proposal.loanAmount * _proposal.interestRate * timeElapsed) / (oneYear * 100);

        return interest;
    }

    // repayLoan is payment of money from borrower to lender
    function repayLoan(uint256 _proposalId) external payable {

        LoanProposal storage  proposal = proposals[_proposalId];

        uint256 repaymentAmount = proposal.loanAmount +  calculateInterest(proposal) ;
  
        // simple case where we require the borrower to pay off one time 
        require(msg.value >= repaymentAmount, "Not enough funds");
        require(proposal.isAccepted && !proposal.isPaidBack, "Loan not accepted or already paid");
        
        // this is where we do the repayment to the lender 
        payable(proposal.lender).transfer(msg.value);

        // the borrower should get back the NFT from contract
        IERC721(proposal.nftContractAddress).transferFrom(address(this), proposal.borrower, proposal.nftTokenId);

        proposal.isPaidBack = true;
    }


    // retractProposal retrieves the NFT from contract and deletes the proposal 
    function retractProposal(uint256 _proposalId) external {
        LoanProposal storage proposal = proposals[_proposalId];

        require(msg.sender == proposal.borrower, "Not the borrower");
        require(!proposal.isAccepted, "Proposal already accepted");

        IERC721(proposal.nftContractAddress).transferFrom(address(this), proposal.borrower, proposal.nftTokenId);
        delete proposals[_proposalId];
    }

    // actOnLoanDateExpiry gives back NFT to lender
    function actOnLoanDateExpiry(uint256 _proposalId) public {
        LoanProposal storage proposal = proposals[_proposalId];

        // Check if its still due 
        require(!proposal.isPaidBack, "Loan is already paid back");

        if (block.timestamp > proposal.dueDate) 
            return;
        
        // Transfer NFT to the lender
        IERC721 nftContract = IERC721(proposal.nftContractAddress);
        nftContract.transferFrom(address(this), proposal.lender, proposal.nftTokenId);

        proposal.isPaidBack = true; //
    }

}

