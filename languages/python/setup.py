from setuptools import setup, find_packages

setup(
    name='your_package_name',
    version='0.1',
    packages=find_packages(),
    include_package_data=True,
    package_data={
        '': ['*.json'],  # Assuming JSON files are in the root or specified directories
    },
)