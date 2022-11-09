FROM willhallonline/ansible
RUN apk --update add --no-cache openssh                                     \
  && sed -i s/#PermitRootLogin.*/PermitRootLogin\ yes/ /etc/ssh/sshd_config \
  && echo "root:1234" | chpasswd                                            \
  && rm -rf /var/cache/apk/*

RUN sed -ie 's/#Port 22/Port 22/g' /etc/ssh/sshd_config
RUN /usr/bin/ssh-keygen -A
RUN ssh-keygen -t rsa -b 4096 -f  /etc/ssh/ssh_host_key

ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile
EXPOSE 22

WORKDIR DeployerData
ADD . .

ENTRYPOINT ["ansible-playbook", "playbook.yml", "-i", "inventory"]
CMD ["/usr/sbin/sshd", "-D"]
