# mackerel-plugin-reboot-required

This is a custom check plugin for [mackerel.io](https://mackerel.io/) to check `/var/run/reboot-required`.

## Usage

```shell
mackerel-plugin-reboot-required [-f=<reboot-required-file>]
```

## Example of mackerel-agent.conf

```toml
[plugin.metrics.reboot-required]
command = "/path/to/mackerel-plugin-reboot-required"
```

## License

MIT License.
