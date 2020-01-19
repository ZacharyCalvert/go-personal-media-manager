# About

This documents the YML format for the meta-data describing the media under the managed folder of this utility.

# Fields

- extensions: possible file extensions available
- paths: locations the file has been found in from import
- reviewDone: if the media has been manually reviewed yet
- ignore: if the file should be ignored on future operations
- sensitive: if the file should be treated as potentially sensitive, such as weigh loss photos
- destroy: if the file is found, destroy if possible else log for destruction
- tags: tags applied, such as for places, friends, events, time, etc for future sorting and organization
- earliestDate: the time millis since epoch according to the file meta-data that the file was first created

# Original

The original format is 1 object per media file, with the media file labeled as an object based on its sha256 sum

```
252F2D6C07AE83BCD39C055288044B7624416A835C7F3856042051D80F35E5B8:
  sha256: 252F2D6C07AE83BCD39C055288044B7624416A835C7F3856042051D80F35E5B8
  extensions:
    - jpg
    - JPG
  paths:
    - path1 
    - path2 
  earliestDate: 1304830800000
  reviewDone: true
  ignore: true
```

# Goal
```
files:
  - <sha>
  paths:
    - 
  earliestDate: 
  reviewDone: 
  tags:
    -
    -
  sensitive:  
  ignore:
  destroy: 
```

