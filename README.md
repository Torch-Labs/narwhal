## init

Initializes new project

`narwhal init <flags>`

### Flags

- `-n` - (Required) project name
- `-g` - (Required) Git repo, in the format `git@github.com:<org>/<repo>.git`
- `-r` - (Required) Remote host, ex: rapsberrypi
- `-u` - (Required) Remote user, ex: pi
- `-e` - (Optinal) Path to .env file, default cwd .env (Will be copied using `scp`), If not specified nothing will be copied
- `-r` - (Optinal) Path to id rsa, default ~/.ssh/id_rsa
- `-c` - (Optional) Path to docker compose file, must be within git workspace. If empty looks for `docker-compose.yml` in cwd
- `-d` - (Optinal) Path to `Dockerfile`, must be within git workspace. If empty looks for `Dockerfile` in cwd
- `-c` - (Optinal) Path to `narwhal_config.yml`, if not specified must use flags
