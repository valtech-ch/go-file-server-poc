# File service

## Start

1.  Replace `<file-url>` with actual file URL
2.  `docker build -t file-service --platform linux/amd64 .`
3.  `docker run -e PORT=3000 -p 3000:3000 file-service`

## Further considerations

- What happens on network downtime?
- Look into request and response timeouts
