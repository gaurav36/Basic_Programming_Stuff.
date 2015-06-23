FROM fedora:22
#FROM fedora20-myimage1:gluster
#FROM fedora20-gluster:latest6
MAINTAINER Gaurav Kumar Garg <garg.gaurav52@gmail.com>

RUN yum -y install python-devel python-setuptools gcc deltarpm yum-utils git  \
                   autoconf automake bison dos2unix flex glib2-devel          \
                   libaio-devel libattr-devel libibverbs-devel                \
                   librdmacm-devel libtool libxml2-devel make openssl-devel   \
                   pkgconfig python-devel python-eventlet python-netifaces    \
                   python-paste-deploy python-simplejson python-sphinx        \
                   python-webob pyxattr readline-devel rpm-build gdb dbench   \
                   net-tools systemtap-sdt-devel attr psmisc findutils which  \
                   xfsprogs yajl-devel lvm2-devel e2fsprogs mock nfs-utils    \
                   openssh-server supervisor openssl fuse-libs wget lvm2-devel\
                   fuse-devel cmockery2-devel userspace-rcu-devel sqlite-devel \
                   libacl-devel vim wget strace supervisor perl-Test-Harness \
                   hostname sudo etcd golang iputils

RUN mkdir -p /var/run/sshd
RUN mkdir -p /home/glusterfs /home/rhs-glusterfs
RUN mkdir -p /var/log/supervisor
RUN echo 'root:password' | chpasswd
RUN ssh-keygen -A
ADD makefusedev.sh   /usr/sbin/makefusedev.sh 
ADD supervisord.conf /etc/supervisord.conf

EXPOSE 22 111 245 443 24007 2049 8080 6010 6011 6012 38465 38466 38468 38469 49152 49153 49154 49156 49157 49158 49159 49160 49161 49162

CMD ["/usr/bin/supervisord"]
