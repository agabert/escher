# escher
This tool turns pictures, movies and music in directories like this:
`/space/backups-before-20230512/zfs001/space/backups/macbook/space/backups/thinkpad/space/backups/old_thinkpad/space/public/pictures`
into this:
`/space/escher/trunk/aaaabbbb.jpg`.

The result is a flat directory with all your unique files in it.

# how does the deduplication work?
The name of the file is based on the checksum of the content.

## Why is this project called escher?
At one point you will deduplicate escher directories like this:
`/space/backups/zfs002/space/escher/trunk`
into this:
`/space/escher/trunk`.

## Why does this not use sqlite?
Because i am not an Apple employee and this is not an intelligent PhotoLibrary.
Keep it simple.