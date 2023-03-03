FROM puppet/puppet-bolt

WORKDIR /installer

RUN apt-get update && apt-get install -y \
  git \
  && rm -rf /var/lib/apt/lists/*

RUN bolt project init installer --modules puppetlabs-peadm

ENTRYPOINT ["bolt", "plan", "run", "peadm::install",  "--params", "@/params.json"]