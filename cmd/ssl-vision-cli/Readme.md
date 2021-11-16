# ssl-vision-cli

The ssl-vision-cli listens for ssl-vision messages and prints them to the console as JSON.

The output can be processed in Bash with `jq`, for example:

```shell
# Print t_capture only
ssl-vision-cli -noGeometry | jq '.t_capture'

# Print t_capture with ball position in csv format
ssl-vision-cli -noGeometry | jq -r '[.t_capture, .balls[0].x, .balls[0].y] | @csv'
```
