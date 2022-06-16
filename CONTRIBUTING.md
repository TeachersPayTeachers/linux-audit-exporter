Contribution Guidelines
=======================

Thank you for contributing! :smile:

 1. Get the requirements:

    * Golang 1.14+
    * Make

 2. Fork and clone repository
    
    ```sh
    git clone git@github.com:<org-or-user>/linux-audit-exporter
    cd linux-audit-exporter
    ```

 3. Create new branch

    ```sh
    git checkout -b <new-branch>
    ```

 4. Build and test your changes

    ```sh
    make build
    make test
    ```

 5. Commit your changes following [conventional commit](https://www.conventionalcommits.org/en/v1.0.0/)

 6. Optionally, run CI tasks locally

    ```sh
    make ci
    ```

 7. Push, and submit pull request

 8. Wait for CI (GitHub actions) to pass

Once the pull request has been reviewed by a Teachers Pay Teachers employee,
and has passed CI, a Teachers Pay Teachers will merge the pull request to the
main branch. At that point, an internal process will automatically create a new
GitHub release, Docker image, and Helm chart if there are any new `fix:` or
`feat:` commits.
