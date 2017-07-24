# tellanch

tellanch - Tell a current branch name of Git using SSH.

## INSTALLATION

```
go get github.com/hiromisuzuki/tellanch
```

## HOW USE

Create config file: `~/.tellanch.yaml`

```yaml
hosts:
  host1:
    user: [SSH username]
    address: [SSH address]
    port: [SSH port|default:22]
    key: [SSH key address]
    path:
      - [Git project path in remote]
  host2:
    ...

      

```

```
$ go run main.go get
```

## AUTHOR

hsuzuki [http://hsuzuki.hatenablog.com/](http://hsuzuki.hatenablog.com/)
