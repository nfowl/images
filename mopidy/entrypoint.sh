#!/bin/sh
# Trigger local library scan
# Hack to enmable scanning on image run
# mopidy local scan

# Switch to the main container `CMD`.
exec "$@"
