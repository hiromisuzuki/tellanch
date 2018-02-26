# tellanch

tellanch - Tell a current branch name of Git, using SSH.

## INSTALLATION

```
go get github.com/hiromisuzuki/tellanch
cd tellanch
go install
```

## HOW TO USE

Create config file: `~/.tellanch.yaml`

```yaml
hosts:
  host1:
    user: [SSH username]
    address: [SSH address]
    port: [SSH port|default:22]
    key: [SSH key address]
    path:
      - [Git project path]
  host2:
    ...

      

```

Run `hosts` command

```
$ tellanch hosts

$ Host 1
    User: [SSH username]
    Address: [SSH address]:[SSH port|default:22]
    Key: [SSH key address]
    Path: [Git project path1]
  Host 2
    User: [SSH username]
    Address: [SSH address]:[SSH port|default:22]
    Key: [SSH key address]
    Path: [Git project path1], [Git project path2]...

```

Run `get` command

```
$ tellanch get

Using config file: /path/to/.tellanch.yaml
[hosts.host1.path1] ~ 
ref: refs/heads/develop
~ 
[hosts.host1.path2] ~ 
ref: refs/heads/feature/foo-feature
~
[hosts.host1.path3] ~ 
ref: refs/heads/feature/bar-feature

```

## AUTHOR

hsuzuki [https://hsuzuki.hatenablog.com/](https://hsuzuki.hatenablog.com/)
