"""
Does the following:
1. Inits git if used
2. Deletes dockerfiles if not going to be used
3. Configures GOPATH and installs dependencies
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
    for filename in ['Dockerfile', '.dockerignore']:
        os.remove(os.path.join(
            PROJECT_DIRECTORY, filename
        ))

def setup_gopath():
    """
    Creates GOPATH structure and moves project files
    """
    thismodule.FINAL_DIRECTORY = os.path.abspath(os.path.join(os.getcwd(), '..', '{{ cookiecutter.gopath }}', 'src', '{{ cookiecutter.pkg }}'))
    logger.info('Setting up GOPATH structure, using package path {}'.format(FINAL_DIRECTORY))

    tmp = tempfile.mkdtemp()
    shutil.move(PROJECT_DIRECTORY, tmp)

    # setup go src dir
    src = os.path.join(PROJECT_DIRECTORY, 'src')
    os.makedirs(src)
    # setup pkg dirs
    pkg = os.path.join(src, '{{ cookiecutter.pkg }}')
    os.makedirs(pkg)

    copy_and_overwrite(os.path.join(tmp, '{{ cookiecutter.name }}'), FINAL_DIRECTORY)
    shutil.rmtree(tmp)

    os.environ['GOPATH'] = PROJECT_DIRECTORY
    os.environ['PATH'] = os.environ['PATH'] + os.path.join(PROJECT_DIRECTORY, 'bin')

def install_deps():
    """
    Installs dependencies with dep
    """

    logger.info('Installing dependencies...')
    dep = Popen(['dep', 'ensure'], cwd=FINAL_DIRECTORY)
    dep.wait()

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
    remove_file('.gitlab-ci.yml')

# 3. Setup GOPATH
if '{{ cookiecutter.pkg }}'.lower() != 'n':
    setup_gopath()

# 4. Install deps
if '{{ cookiecutter.install }}'.lower() != 'n':
    install_deps()

# 5. Initialize Git (should be run after all file have been modified or deleted)
if '{{ cookiecutter.origin }}'.lower() != 'n':
    init_git()
else:
    remove_file('.gitignore')

logger.info('Your project is ready to go, to start working:')
if '{{ cookiecutter.install }}'.lower() != 'n':
    logger.info('`cd {0}/src/{1}`'.format(PROJECT_DIRECTORY, '{{ cookiecutter.pkg }}'))
    logger.info('`export GOPATH={}`'.format(PROJECT_DIRECTORY))
else:
    logger.info('Manually configure your GOPATH')
    logger.info('Install deps: `dep ensure`')
