from setuptools import find_packages, setup

# yapf: disable
setup(
    name="sfxpython",
    version="0.4.2",
    author="SignalFx, Inc",
    description="Python packages used by the Python extension mechanism of the SignalFx Smart Agent",
    url="https://github.com/signalfx/signalfx-agent",
    packages=find_packages(),
    install_requires=[
        'ujson==4.2.0',
    ],
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: Apache Software License",
    ]
)
