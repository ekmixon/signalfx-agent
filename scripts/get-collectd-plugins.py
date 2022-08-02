#!/usr/bin/env python3

import contextlib
import os
import shutil
import subprocess
import sys
import tarfile
import urllib.request

import yaml

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))
TARGET_DIR = os.path.join("/", "usr", "share", "collectd") if len(sys.argv) < 2 else sys.argv[1]
PYTHON_EXECUTABLE = sys.executable or "python"

with open(os.path.join(SCRIPT_DIR, "..", "collectd-plugins.yaml"), "r") as f:
    PLUGINS = yaml.safe_load(f)

for p in PLUGINS:
    plugin_name = p.get("name")
    version = p.get("version")
    repo = p.get("repo")
    url = "https://github.com/{repo}/archive/{version}.tar.gz".format(repo=repo, version=version)

    print(
        """Bundling...
plugin:  {p}
version: {v}
repo:    {r}
url:     {u}""".format(
            p=plugin_name, v=version, r=repo, u=url
        )
    )

    with contextlib.closing(urllib.request.urlopen(url)) as stream:
        with tarfile.open(fileobj=stream, mode="r|gz") as tar_archive:
            tar_archive.extractall(TARGET_DIR)
            plugin_dir = os.path.join(TARGET_DIR, plugin_name)
            os.rename(os.path.join(TARGET_DIR, tar_archive.getnames()[0]), plugin_dir)

    # install pip deps
    for package in p.get("pip_packages", []):
        subprocess.check_call([PYTHON_EXECUTABLE, "-m", "pip", "install", "-qq", "--no-warn-script-location", package])

    requirements_file = os.path.join(plugin_dir, "requirements.txt")
    if os.path.isfile(requirements_file):
        subprocess.check_call(
            [PYTHON_EXECUTABLE, "-m", "pip", "install", "-qq", "--no-warn-script-location", "-r", requirements_file]
        )

    # remove unecessary things
    for elem in p.get("can_remove", []):
        shutil.rmtree(os.path.join(plugin_dir, elem))
