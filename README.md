# go-filetype (WIP)

## Supported file formats

### Image

| Extension | MIME                     | Description                                                   | Condition | Comment |
| --------- | ------------------------ | ------------------------------------------------------------- | --------- | ------- |
| bmp       | image/bmp                | A bitmap format used mostly in the Windows world              |           |         |
| gif       | image/gif                | Image file encoded in the Graphics Interchange Format         |           |         |
| ico       | image/vnd.microsoft.icon | Computer icon encoded in ICO file format                      |           |         |
| jp2       | image/jp2                | JPEG 2000 format                                              |           |         |
| jpa       | image/jpa                | JPEG 2000 format                                              |           |         |
| jpg       | image/jpeg               | JPEG using JFIF/EXIF/SPIFF                                    |           |         |
| jpm       | image/jpm                | JPEG 2000 format                                              |           |         |
| jpx       | image/jpx                | JPEG 2000 format                                              |           |         |
| mj2       | image/mj2                | JPEG 2000 format                                              |           |         |
| png       | image/png                | Image encoded in the Portable Network Graphics format         |           |         |
| psd       | vnd.adobe.photoshop      | Photoshop Document file, Adobe Photoshop's native file format |           |         |
| tiff      | image/tiff               | Tagged Image File Format                                      |           |         |
| webp      | image/webp               | Google WebP image file                                        |           |         |
|           |                          |                                                               |           |         |

### Video

| Extension | MIME             | Description                                                           | Condition | Comment |
| --------- | ---------------- | --------------------------------------------------------------------- | --------- | ------- |
| 3gp       | video/3gpp       | 3rd Generation Partnership Project 3GPP multimedia file               |           |         |
| avi       | video/x-msvideo  | Audio Video Interleave video format                                   |           |         |
| flv       | video/x-flv      | Flash Video file                                                      |           |         |
| m4v       | video/x-m4v      | ISO Media, Apple iTunes Video (.M4V) Video                            |           |         |
| mkv       | video/x-matroska | Matroska media container                                              |           |         |
| mov       | video/quicktime  | ISO Media, Apple QuickTime movie, Apple QuickTime (.MOV/QT)           |           |         |
| mpg       | video/mpeg       | MPEG video file / MPEG Program Stream / MPEG-1 video and MPEG-2 video |           |         |
| ogm       | video/ogg        | OGG Media Stream file                                                 |           |         |
| ogv       | video/ogg        | Ogg Video file                                                        |           |         |
| webm      | video/webm       | Matroska media container (webm)                                       |           |         |
| wmv       | video/x-ms-wmv   | Microsoft Windows Media Audio/Video File                              |           |         |
|           |                  |                                                                       |           |         |

### Audio

| Extension | MIME            | Description                                            | Condition | Comment |
| --------- | --------------- | ------------------------------------------------------ | --------- | ------- |
| aiff      | audio/x-aiff    | Audio Interchange File Format                          |           |         |
| flac      | audio/x-flac    | Free Lossless Audio Codec                              |           |         |
| m4a       | audio/mp4       | ISO Media, Apple iTunes Audio encoded with AAC or ALAC |           |         |
| m4p       | audio/m4p       | ISO Media, Apple iTunes Audio                          |           |         |
| mid       | audio/midi      | MIDI sound file                                        |           |         |
| oga       | audio/ogg       | An open source media container format (Audio)          |           |         |
| ogg       | audio/ogg       | An open source media container format (Audio)          |           |         |
| ogx       | application/ogg | An open source media container format                  |           |         |
| opus      | audio/opus      | An open source media container format (Audio)          |           |         |
| spx       | audio/ogg       | An open source media container format (Audio)          |           |         |
| wav       | audio/x-wav     | Waveform Audio File Format                             |           |         |

### Archives

| Extension | MIME | Description | Condition | Comment |
| --------- | ---- | ----------- | --------- | ------- |
| Z         |      |             |           |         |
| odt       |      |             |           |         |
| odp       |      |             |           |         |
| xpi       |      |             |           |         |
| zip       |      |             |           |         |
| rar       |      |             |           |         |
| tar       |      |             |           |         |
| 7z        |      |             |           |         |

### Office and documents

| Extension | MIME            | Description              | Condition | Comment |
| --------- | --------------- | ------------------------ | --------- | ------- |
| doc       |                 |                          |           |         |
| ppt       |                 |                          |           |         |
| xls       |                 |                          |           |         |
| docx      |                 |                          |           |         |
| pptx      |                 |                          |           |         |
| xslx      |                 |                          |           |         |
| epub      |                 |                          |           |         |
| pdf       | application/pdf | Portable Document Format |           |         |

### Other

| Extension | MIME                                  | Description                                | Condition | Comment |
| --------- | ------------------------------------- | ------------------------------------------ | --------- | ------- |
| apk       | application/apk-archive               | Android package                            |           | WIP     |
| class     | application/x-java-class              | Java class file                            |           |         |
| deb       | application/vnd.debian.binary-package | Linux deb file                             |           |         |
| dll       | application/x-msdownload              | Dynamic-link library for Microsoft Windows |           |         |
| exe       | application/x-msdownload              | Executable file for Microsoft Windows      |           |         |
| iso       | application/x-iso9660-image           | ISO9660 CD/DVD image file                  |           |         |
| jar       | application/java-archive              | Java archive                               |           |         |
| lnk       | application/x-ms-shortcut             | Windows shell link (shortcut) file         |           |         |
| msi       | application/x-msi                     | Microsoft Software Installer               |           |         |
| vdi       | application/x-virtualbox-vdi          | Oracle VirtualBox Image                    |           |         |

## TODO

- ELF
- .lib Could not find much about the format, might use `ar` file format: https://en.wikipedia.org/wiki/Ar_(Unix)#File_format_details
- .so
- .ad1
- .vmem (53) FF 00 F0 53 FF 00 F0?
- .vmss/.vmsn
- .vmdk
- .scr
- encrypted zip?
- .zlib
- .msg
- .gz
- database formats
