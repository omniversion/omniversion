# Integration tests

Test the entire pipeline: Ansible roles, CLI and Python module. 

## How to run

1. Create containers
```shell
docker-compose up base-instance && docker-compose up
```

2. Create & activate virtual environment
```shell
python3 -m venv env && source env/bin/activate
```

3. Install requirements
```shell
pip install -r test-requirements.txt
```

4. Run molecule tests
```shell
molecule test
```
