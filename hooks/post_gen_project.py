"""
Does the following:
1. Inits git if used
2. Deletes Dockerfile if not going to be used
3. Deletes CI config based on selection
"""
from __future__ import print_function
import logging
import os
import shutil
import sys
import tempfile
from subprocess import Popen


logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger('gocookiecutter')
thismodule = sys.modules[__name__]

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)
FINAL_DIRECTORY = PROJECT_DIRECTORY

def remove_file(filename):
    """
    generic remove file from project dir
    """
    fullpath = os.path.join(PROJECT_DIRECTORY, filename)
    if os.path.exists(fullpath):
        if os.path.isdir(fullpath):
            shutil.rmtree(fullpath)
            return
        os.remove(fullpath)

def copy_and_overwrite(src, dst):
    """
    Removes existing directory from `dst` and overwrites with `src`
    """
    if os.path.exists(dst):
        shutil.rmtree(dst)
    shutil.copytree(src, dst)

def remove_docker_files():
    """
    Removes files needed for docker if it isn't going to be used
    """
    for filename in ['./build/package', '.dockerignore']:
        remove_file(filename)

def init_git():
    """
    Initialises git on the new project folder
    """
    GIT_COMMANDS = [
        ['git', 'init'],
        ['git', 'add', '.'],
        ['git', 'commit', '-aq', '-m', 'Initial Commit'],
        ['git', 'remote', 'add', 'origin', '{{ cookiecutter.origin }}']
    ]

    for command in GIT_COMMANDS:
        git = Popen(command, cwd=FINAL_DIRECTORY)
        git.wait()

# 1. Remove Dockerfiles if docker is not going to be used
if '{{ cookiecutter.image }}'.lower() == 'n':
    remove_docker_files()

# 2. Remove unused CI choice
if '{{ cookiecutter.use_ci}}'.lower() == 'gitlab':
    # do nothing
    pass
else:
    remove_file('./build/ci/.gitlab-ci.yml')

# 3. Initialize Git (should be run after all files have been modified or deleted)
if '{{ cookiecutter.origin }}'.lower() != 'n':
    init_git()
else:
    remove_file('.gitignore')

logger.info('Your project is ready to go, to start working:')
logger.info('`cd {0}`'.format(PROJECT_DIRECTORY))
