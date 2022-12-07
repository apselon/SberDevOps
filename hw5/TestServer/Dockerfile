FROM jenkins/jenkins:lts

USER root 
RUN apt update && apt install -y python python3-pytest python3-pip && pip3 install pytest-cov allure-pytest
