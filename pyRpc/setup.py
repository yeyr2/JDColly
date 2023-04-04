from setuptools import setup, find_packages

setup(
    name="pyRpc",
    version="0.1",
    packages=find_packages(),
    entry_points={
        "console_scripts": [
            "pyRpc = pyRpc.__main__:main",
        ]
    },
)
