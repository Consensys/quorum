[![Documentation Status](https://readthedocs.org/projects/goquorum/badge/?version=latest)](http://docs.goquorum.com/en/latest/?badge=latest)

# Quorum documentation

New Quorum documentation is now published on https://docs.goquorum.com/

## How to contribute

Quorum documentation files are written in Markdown and configured with a
YAML configuration file from [mkdocs](https://www.mkdocs.org/)

The documentation site uses [Material theme](https://squidfunk.github.io/mkdocs-material/)
which has been configured with number of theme [extensions](https://squidfunk.github.io/mkdocs-material/extensions/admonition/)
to enhance documenting experience 

To contribute, here is 3 simple steps

- Use Python to install `mkdocs` and related dependencies
    ```bash
    pip install -r requirements.txt
    ```
- Add/Modify desired documentation files and preview the site
    ```bash
    mkdocs serve
    ```
- Commit and raise PR against master