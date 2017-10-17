/** @file
 *****************************************************************************

 Implementation of interfaces for initializing MNT4.

 See mnt4_init.hpp .

 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#include "algebra/curves/mnt/mnt4/mnt4_init.hpp"
#include "algebra/curves/mnt/mnt4/mnt4_g1.hpp"
#include "algebra/curves/mnt/mnt4/mnt4_g2.hpp"

namespace libsnark {

// bigint<mnt4_r_limbs> mnt4_modulus_r = mnt46_modulus_A;
// bigint<mnt4_q_limbs> mnt4_modulus_q = mnt46_modulus_B;

mnt4_Fq2 mnt4_twist;
mnt4_Fq2 mnt4_twist_coeff_a;
mnt4_Fq2 mnt4_twist_coeff_b;
mnt4_Fq mnt4_twist_mul_by_a_c0;
mnt4_Fq mnt4_twist_mul_by_a_c1;
mnt4_Fq mnt4_twist_mul_by_b_c0;
mnt4_Fq mnt4_twist_mul_by_b_c1;
mnt4_Fq mnt4_twist_mul_by_q_X;
mnt4_Fq mnt4_twist_mul_by_q_Y;

bigint<mnt4_q_limbs> mnt4_ate_loop_count;
bool mnt4_ate_is_loop_count_neg;
bigint<4*mnt4_q_limbs> mnt4_final_exponent;
bigint<mnt4_q_limbs> mnt4_final_exponent_last_chunk_abs_of_w0;
bool mnt4_final_exponent_last_chunk_is_w0_neg;
bigint<mnt4_q_limbs> mnt4_final_exponent_last_chunk_w1;

void init_mnt4_params()
{
    typedef bigint<mnt4_r_limbs> bigint_r;
    typedef bigint<mnt4_q_limbs> bigint_q;

    assert(sizeof(mp_limb_t) == 8 || sizeof(mp_limb_t) == 4); // Montgomery assumes this

    /* parameters for scalar field Fr */
    mnt4_modulus_r = bigint_r("475922286169261325753349249653048451545124878552823515553267735739164647307408490559963137");
    assert(mnt4_Fr::modulus_is_valid());
    if (sizeof(mp_limb_t) == 8)
    {
        mnt4_Fr::Rsquared = bigint_r("163983144722506446826715124368972380525894397127205577781234305496325861831001705438796139");
        mnt4_Fr::Rcubed = bigint_r("207236281459091063710247635236340312578688659363066707916716212805695955118593239854980171");
        mnt4_Fr::inv = 0xbb4334a3ffffffff;
    }
    if (sizeof(mp_limb_t) == 4)
    {
        mnt4_Fr::Rsquared = bigint_r("163983144722506446826715124368972380525894397127205577781234305496325861831001705438796139");
        mnt4_Fr::Rcubed = bigint_r("207236281459091063710247635236340312578688659363066707916716212805695955118593239854980171");
        mnt4_Fr::inv = 0xffffffff;
    }
    mnt4_Fr::num_bits = 298;
    mnt4_Fr::euler = bigint_r("237961143084630662876674624826524225772562439276411757776633867869582323653704245279981568");
    mnt4_Fr::s = 34;
    mnt4_Fr::t = bigint_r("27702323054502562488973446286577291993024111641153199339359284829066871159442729");
    mnt4_Fr::t_minus_1_over_2 = bigint_r("13851161527251281244486723143288645996512055820576599669679642414533435579721364");
    mnt4_Fr::multiplicative_generator = mnt4_Fr("10");
    mnt4_Fr::root_of_unity = mnt4_Fr("120638817826913173458768829485690099845377008030891618010109772937363554409782252579816313");
    mnt4_Fr::nqr = mnt4_Fr("5");
    mnt4_Fr::nqr_to_t = mnt4_Fr("406220604243090401056429458730298145937262552508985450684842547562990900634752279902740880");

    /* parameters for base field Fq */
    mnt4_modulus_q = bigint_q("475922286169261325753349249653048451545124879242694725395555128576210262817955800483758081");
    assert(mnt4_Fq::modulus_is_valid());
    if (sizeof(mp_limb_t) == 8)
    {
        mnt4_Fq::Rsquared = bigint_q("273000478523237720910981655601160860640083126627235719712980612296263966512828033847775776");
        mnt4_Fq::Rcubed = bigint_q("427298980065529822574935274648041073124704261331681436071990730954930769758106792920349077");
        mnt4_Fq::inv = 0xb071a1b67165ffff;
    }
    if (sizeof(mp_limb_t) == 4)
    {
        mnt4_Fq::Rsquared = bigint_q("273000478523237720910981655601160860640083126627235719712980612296263966512828033847775776");
        mnt4_Fq::Rcubed = bigint_q("427298980065529822574935274648041073124704261331681436071990730954930769758106792920349077");
        mnt4_Fq::inv = 0x7165ffff;
    }
    mnt4_Fq::num_bits = 298;
    mnt4_Fq::euler = bigint_q("237961143084630662876674624826524225772562439621347362697777564288105131408977900241879040");
    mnt4_Fq::s = 17;
    mnt4_Fq::t = bigint_q("3630998887399759870554727551674258816109656366292531779446068791017229177993437198515");
    mnt4_Fq::t_minus_1_over_2 = bigint_q("1815499443699879935277363775837129408054828183146265889723034395508614588996718599257");
    mnt4_Fq::multiplicative_generator = mnt4_Fq("17");
    mnt4_Fq::root_of_unity = mnt4_Fq("264706250571800080758069302369654305530125675521263976034054878017580902343339784464690243");
    mnt4_Fq::nqr = mnt4_Fq("17");
    mnt4_Fq::nqr_to_t = mnt4_Fq("264706250571800080758069302369654305530125675521263976034054878017580902343339784464690243");

    /* parameters for twist field Fq2 */
    mnt4_Fq2::euler = bigint<2*mnt4_q_limbs>("113251011236288135098249345249154230895914381858788918106847214243419142422924133497460817468249854833067260038985710370091920860837014281886963086681184370139950267830740466401280");
    mnt4_Fq2::s = 18;
    mnt4_Fq2::t = bigint<2*mnt4_q_limbs>("864036645784668999467844736092790457885088972921668381552484239528039111503022258739172496553419912972009735404859240494475714575477709059806542104196047745818712370534824115");
    mnt4_Fq2::t_minus_1_over_2 = bigint<2*mnt4_q_limbs>("432018322892334499733922368046395228942544486460834190776242119764019555751511129369586248276709956486004867702429620247237857287738854529903271052098023872909356185267412057");
    mnt4_Fq2::non_residue = mnt4_Fq("17");
    mnt4_Fq2::nqr = mnt4_Fq2(mnt4_Fq("8"),mnt4_Fq("1"));
    mnt4_Fq2::nqr_to_t = mnt4_Fq2(mnt4_Fq("0"),mnt4_Fq("29402818985595053196743631544512156561638230562612542604956687802791427330205135130967658"));
    mnt4_Fq2::Frobenius_coeffs_c1[0] = mnt4_Fq("1");
    mnt4_Fq2::Frobenius_coeffs_c1[1] = mnt4_Fq("475922286169261325753349249653048451545124879242694725395555128576210262817955800483758080");

    /* parameters for Fq4 */
    mnt4_Fq4::non_residue = mnt4_Fq("17");
    mnt4_Fq4::Frobenius_coeffs_c1[0] = mnt4_Fq("1");
    mnt4_Fq4::Frobenius_coeffs_c1[1] = mnt4_Fq("7684163245453501615621351552473337069301082060976805004625011694147890954040864167002308");
    mnt4_Fq4::Frobenius_coeffs_c1[2] = mnt4_Fq("475922286169261325753349249653048451545124879242694725395555128576210262817955800483758080");
    mnt4_Fq4::Frobenius_coeffs_c1[3] = mnt4_Fq("468238122923807824137727898100575114475823797181717920390930116882062371863914936316755773");

    /* choice of short Weierstrass curve and its twist */
    mnt4_G1::coeff_a = mnt4_Fq("2");
    mnt4_G1::coeff_b = mnt4_Fq("423894536526684178289416011533888240029318103673896002803341544124054745019340795360841685");
    mnt4_twist = mnt4_Fq2(mnt4_Fq::zero(), mnt4_Fq::one());
    mnt4_twist_coeff_a = mnt4_Fq2(mnt4_G1::coeff_a * mnt4_Fq2::non_residue, mnt4_Fq::zero());
    mnt4_twist_coeff_b = mnt4_Fq2(mnt4_Fq::zero(), mnt4_G1::coeff_b * mnt4_Fq2::non_residue);
    mnt4_G2::twist = mnt4_twist;
    mnt4_G2::coeff_a = mnt4_twist_coeff_a;
    mnt4_G2::coeff_b = mnt4_twist_coeff_b;
    mnt4_twist_mul_by_a_c0 = mnt4_G1::coeff_a * mnt4_Fq2::non_residue;
    mnt4_twist_mul_by_a_c1 = mnt4_G1::coeff_a * mnt4_Fq2::non_residue;
    mnt4_twist_mul_by_b_c0 = mnt4_G1::coeff_b * mnt4_Fq2::non_residue.squared();
    mnt4_twist_mul_by_b_c1 = mnt4_G1::coeff_b * mnt4_Fq2::non_residue;
    mnt4_twist_mul_by_q_X = mnt4_Fq("475922286169261325753349249653048451545124879242694725395555128576210262817955800483758080");
    mnt4_twist_mul_by_q_Y = mnt4_Fq("7684163245453501615621351552473337069301082060976805004625011694147890954040864167002308");

    /* choice of group G1 */
    mnt4_G1::G1_zero = mnt4_G1(mnt4_Fq::zero(),
                               mnt4_Fq::one(),
                               mnt4_Fq::zero());


    mnt4_G1::G1_one = mnt4_G1(mnt4_Fq("60760244141852568949126569781626075788424196370144486719385562369396875346601926534016838"),
                              mnt4_Fq("363732850702582978263902770815145784459747722357071843971107674179038674942891694705904306"),
                              mnt4_Fq::one());

    mnt4_G1::wnaf_window_table.resize(0);
    mnt4_G1::wnaf_window_table.push_back(11);
    mnt4_G1::wnaf_window_table.push_back(24);
    mnt4_G1::wnaf_window_table.push_back(60);
    mnt4_G1::wnaf_window_table.push_back(127);

    mnt4_G1::fixed_base_exp_window_table.resize(0);
    // window 1 is unbeaten in [-inf, 5.09]
    mnt4_G1::fixed_base_exp_window_table.push_back(1);
    // window 2 is unbeaten in [5.09, 9.64]
    mnt4_G1::fixed_base_exp_window_table.push_back(5);
    // window 3 is unbeaten in [9.64, 24.79]
    mnt4_G1::fixed_base_exp_window_table.push_back(10);
    // window 4 is unbeaten in [24.79, 60.29]
    mnt4_G1::fixed_base_exp_window_table.push_back(25);
    // window 5 is unbeaten in [60.29, 144.37]
    mnt4_G1::fixed_base_exp_window_table.push_back(60);
    // window 6 is unbeaten in [144.37, 344.90]
    mnt4_G1::fixed_base_exp_window_table.push_back(144);
    // window 7 is unbeaten in [344.90, 855.00]
    mnt4_G1::fixed_base_exp_window_table.push_back(345);
    // window 8 is unbeaten in [855.00, 1804.62]
    mnt4_G1::fixed_base_exp_window_table.push_back(855);
    // window 9 is unbeaten in [1804.62, 3912.30]
    mnt4_G1::fixed_base_exp_window_table.push_back(1805);
    // window 10 is unbeaten in [3912.30, 11264.50]
    mnt4_G1::fixed_base_exp_window_table.push_back(3912);
    // window 11 is unbeaten in [11264.50, 27897.51]
    mnt4_G1::fixed_base_exp_window_table.push_back(11265);
    // window 12 is unbeaten in [27897.51, 57596.79]
    mnt4_G1::fixed_base_exp_window_table.push_back(27898);
    // window 13 is unbeaten in [57596.79, 145298.71]
    mnt4_G1::fixed_base_exp_window_table.push_back(57597);
    // window 14 is unbeaten in [145298.71, 157204.59]
    mnt4_G1::fixed_base_exp_window_table.push_back(145299);
    // window 15 is unbeaten in [157204.59, 601600.62]
    mnt4_G1::fixed_base_exp_window_table.push_back(157205);
    // window 16 is unbeaten in [601600.62, 1107377.25]
    mnt4_G1::fixed_base_exp_window_table.push_back(601601);
    // window 17 is unbeaten in [1107377.25, 1789646.95]
    mnt4_G1::fixed_base_exp_window_table.push_back(1107377);
    // window 18 is unbeaten in [1789646.95, 4392626.92]
    mnt4_G1::fixed_base_exp_window_table.push_back(1789647);
    // window 19 is unbeaten in [4392626.92, 8221210.60]
    mnt4_G1::fixed_base_exp_window_table.push_back(4392627);
    // window 20 is unbeaten in [8221210.60, 42363731.19]
    mnt4_G1::fixed_base_exp_window_table.push_back(8221211);
    // window 21 is never the best
    mnt4_G1::fixed_base_exp_window_table.push_back(0);
    // window 22 is unbeaten in [42363731.19, inf]
    mnt4_G1::fixed_base_exp_window_table.push_back(42363731);

    /* choice of group G2 */
    mnt4_G2::G2_zero = mnt4_G2(mnt4_Fq2::zero(),
                               mnt4_Fq2::one(),
                               mnt4_Fq2::zero());

    mnt4_G2::G2_one = mnt4_G2(mnt4_Fq2(mnt4_Fq("438374926219350099854919100077809681842783509163790991847867546339851681564223481322252708"),
                                       mnt4_Fq("37620953615500480110935514360923278605464476459712393277679280819942849043649216370485641")),
                              mnt4_Fq2(mnt4_Fq("37437409008528968268352521034936931842973546441370663118543015118291998305624025037512482"),
                                       mnt4_Fq("424621479598893882672393190337420680597584695892317197646113820787463109735345923009077489")),
                              mnt4_Fq2::one());

    mnt4_G2::wnaf_window_table.resize(0);
    mnt4_G2::wnaf_window_table.push_back(5);
    mnt4_G2::wnaf_window_table.push_back(15);
    mnt4_G2::wnaf_window_table.push_back(39);
    mnt4_G2::wnaf_window_table.push_back(109);

    mnt4_G2::fixed_base_exp_window_table.resize(0);
    // window 1 is unbeaten in [-inf, 4.17]
    mnt4_G2::fixed_base_exp_window_table.push_back(1);
    // window 2 is unbeaten in [4.17, 10.12]
    mnt4_G2::fixed_base_exp_window_table.push_back(4);
    // window 3 is unbeaten in [10.12, 24.65]
    mnt4_G2::fixed_base_exp_window_table.push_back(10);
    // window 4 is unbeaten in [24.65, 60.03]
    mnt4_G2::fixed_base_exp_window_table.push_back(25);
    // window 5 is unbeaten in [60.03, 143.16]
    mnt4_G2::fixed_base_exp_window_table.push_back(60);
    // window 6 is unbeaten in [143.16, 344.73]
    mnt4_G2::fixed_base_exp_window_table.push_back(143);
    // window 7 is unbeaten in [344.73, 821.24]
    mnt4_G2::fixed_base_exp_window_table.push_back(345);
    // window 8 is unbeaten in [821.24, 1793.92]
    mnt4_G2::fixed_base_exp_window_table.push_back(821);
    // window 9 is unbeaten in [1793.92, 3919.59]
    mnt4_G2::fixed_base_exp_window_table.push_back(1794);
    // window 10 is unbeaten in [3919.59, 11301.46]
    mnt4_G2::fixed_base_exp_window_table.push_back(3920);
    // window 11 is unbeaten in [11301.46, 18960.09]
    mnt4_G2::fixed_base_exp_window_table.push_back(11301);
    // window 12 is unbeaten in [18960.09, 44198.62]
    mnt4_G2::fixed_base_exp_window_table.push_back(18960);
    // window 13 is unbeaten in [44198.62, 150799.57]
    mnt4_G2::fixed_base_exp_window_table.push_back(44199);
    // window 14 is never the best
    mnt4_G2::fixed_base_exp_window_table.push_back(0);
    // window 15 is unbeaten in [150799.57, 548694.81]
    mnt4_G2::fixed_base_exp_window_table.push_back(150800);
    // window 16 is unbeaten in [548694.81, 1051769.08]
    mnt4_G2::fixed_base_exp_window_table.push_back(548695);
    // window 17 is unbeaten in [1051769.08, 2023925.59]
    mnt4_G2::fixed_base_exp_window_table.push_back(1051769);
    // window 18 is unbeaten in [2023925.59, 3787108.68]
    mnt4_G2::fixed_base_exp_window_table.push_back(2023926);
    // window 19 is unbeaten in [3787108.68, 7107480.30]
    mnt4_G2::fixed_base_exp_window_table.push_back(3787109);
    // window 20 is unbeaten in [7107480.30, 38760027.14]
    mnt4_G2::fixed_base_exp_window_table.push_back(7107480);
    // window 21 is never the best
    mnt4_G2::fixed_base_exp_window_table.push_back(0);
    // window 22 is unbeaten in [38760027.14, inf]
    mnt4_G2::fixed_base_exp_window_table.push_back(38760027);

    /* pairing parameters */
    mnt4_ate_loop_count = bigint_q("689871209842287392837045615510547309923794944");
    mnt4_ate_is_loop_count_neg = false;
    mnt4_final_exponent = bigint<4*mnt4_q_limbs>("107797360357109903430794490309592072278927783803031854357910908121903439838772861497177116410825586743089760869945394610511917274977971559062689561855016270594656570874331111995170645233717143416875749097203441437192367065467706065411650403684877366879441766585988546560");
    mnt4_final_exponent_last_chunk_abs_of_w0 = bigint_q("689871209842287392837045615510547309923794945");
    mnt4_final_exponent_last_chunk_is_w0_neg = false;
    mnt4_final_exponent_last_chunk_w1 = bigint_q("1");
}

} // libsnark
