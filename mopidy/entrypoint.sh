#!/bin/sh
# Trigger local library scan
# Hack to enmable scanning via pod recreation
mopidy local scan

# Switch to the main container `CMD`.
exec "$@"
