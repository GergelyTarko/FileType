# go-filetype (WIP)

## Apparently Microsoft Defender might flag the build as Trojan

## Supported file formats

### Image

| Extension | MIME                     | Description                                                   | Comment |
| --------- | ------------------------ | ------------------------------------------------------------- | ------- |
| bmp       | image/bmp                | A bitmap format used mostly in the Windows world              |         |
| gif       | image/gif                | Image file encoded in the Graphics Interchange Format         |         |
| ico       | image/vnd.microsoft.icon | Computer icon encoded in ICO file format                      |         |
| jp2       | image/jp2                | JPEG 2000 format                                              |         |
| jpa       | image/jpa                | JPEG 2000 format                                              |         |
| jpg       | image/jpeg               | JPEG using JFIF/EXIF/SPIFF                                    |         |
| jpm       | image/jpm                | JPEG 2000 format                                              |         |
| jpx       | image/jpx                | JPEG 2000 format                                              |         |
| mj2       | image/mj2                | JPEG 2000 format                                              |         |
| png       | image/png                | Image encoded in the Portable Network Graphics format         |         |
| psd       | vnd.adobe.photoshop      | Photoshop Document file, Adobe Photoshop's native file format |         |
| tiff      | image/tiff               | Tagged Image File Format                                      |         |
| webp      | image/webp               | Google WebP image file                                        |         |

### Video

| Extension | MIME             | Description                                                           | Comment |
| --------- | ---------------- | --------------------------------------------------------------------- | ------- |
| 3gp       | video/3gpp       | 3rd Generation Partnership Project 3GPP multimedia file               |         |
| avi       | video/x-msvideo  | Audio Video Interleave video format                                   |         |
| flv       | video/x-flv      | Flash Video file                                                      |         |
| m4v       | video/x-m4v      | ISO Media, Apple iTunes Video (.M4V) Video                            |         |
| mkv       | video/x-matroska | Matroska media container                                              |         |
| mov       | video/quicktime  | ISO Media, Apple QuickTime movie, Apple QuickTime (.MOV/QT)           |         |
| mpg       | video/mpeg       | MPEG video file / MPEG Program Stream / MPEG-1 video and MPEG-2 video |         |
| ogm       | video/ogg        | OGG Media Stream file                                                 |         |
| ogv       | video/ogg        | Ogg Video file                                                        |         |
| webm      | video/webm       | Matroska media container (webm)                                       |         |
| wmv       | video/x-ms-wmv   | Microsoft Windows Media Audio/Video File                              |         |

### Audio

| Extension | MIME            | Description                                            | Comment |
| --------- | --------------- | ------------------------------------------------------ | ------- |
| aiff      | audio/x-aiff    | Audio Interchange File Format                          |         |
| flac      | audio/x-flac    | Free Lossless Audio Codec                              |         |
| m4a       | audio/mp4       | ISO Media, Apple iTunes Audio encoded with AAC or ALAC |         |
| m4p       | audio/m4p       | ISO Media, Apple iTunes Audio                          |         |
| mid       | audio/midi      | MIDI sound file                                        |         |
| oga       | audio/ogg       | An open source media container format (Audio)          |         |
| ogg       | audio/ogg       | An open source media container format (Audio)          |         |
| ogx       | application/ogg | An open source media container format                  |         |
| opus      | audio/opus      | An open source media container format (Audio)          |         |
| spx       | audio/ogg       | An open source media container format (Audio)          |         |
| wav       | audio/x-wav     | Waveform Audio File Format                             |         |

### Archives

| Extension | MIME                                            | Description                                     | Comment |
| --------- | ----------------------------------------------- | ----------------------------------------------- | ------- |
| 7z        | application/x-7z-compressed                     | 7-Zip File Format                               |         |
| odp       | application/vnd.oasis.opendocument.presentation | OpenOffice Presentation Document                |         |
| ods       | application/vnd.oasis.opendocument.spreadsheet  | OpenDocument Spreadsheet Document               |         |
| odt       | application/vnd.oasis.opendocument.text         | OpenOffice Text Document                        |         |
| rar       | application/vnd.rar                             | Roshal ARchive compressed archive v1.50 onwards |         |
| tar       | application/x-tar                               | tar archive                                     |         |
| xpi       | application/x-xpinstall                         | Mozilla Browser Archive                         |         |
| Z         | application/x-compress                          | LZW/LZH                                         |         |
| zip       | application/zip                                 | PKZIP archive file                              |         |

### Office and documents

| Extension | MIME                                                                      | Description                                          | Comment                                |
| --------- | ------------------------------------------------------------------------- | ---------------------------------------------------- | -------------------------------------- |
| doc       | application/msword                                                        | MS Office Word document file                         | Word 2.0 file might cannot be detected |
| docx      | application/vnd.openxmlformats-officedocument.wordprocessingml.document   | Office Open XML File Format                          |                                        |
| epub      | application/epub+zip                                                      | E-book file format                                   |                                        |
| msg       | application/vnd.ms-outlook                                                | Microsoft Outlook and MS Exchange email message file |                                        |
| pdf       | application/pdf                                                           | Portable Document Format                             |                                        |
| ppt       | application/vnd.ms-powerpoint                                             | MS Office PowerPoint presentation file               |                                        |
| pptx      | application/vnd.openxmlformats-officedocument.presentationml.presentation | Office Open XML File Format                          |                                        |
| xls       | application/vnd.ms-excel                                                  | MS Office Word excel file                            |                                        |
| xslx      | application/vnd.openxmlformats-officedocument.spreadsheetml.sheet         | Office Open XML File Format                          |                                        |

### Other

| Extension          | MIME                                  | Description                                | Comment                                                                              |
| ------------------ | ------------------------------------- | ------------------------------------------ | ------------------------------------------------------------------------------------ |
| ad1                | application/octet-stream              | Forensic Toolkit FTK Imager Image file     |                                                                                      |
| apk                | application/apk-archive               | Android package                            | WIP                                                                                  |
| class              | application/x-java-class              | Java class file                            |                                                                                      |
| deb                | application/vnd.debian.binary-package | Linux deb file                             |                                                                                      |
| dll                | application/x-msdownload              | Dynamic-link library for Microsoft Windows |                                                                                      |
| elf                | application/x-elf                     | Executable and Linkable Format             |                                                                                      |
| evt                | application/octet-stream              | Windows Event Viewer file format           |                                                                                      |
| evtx               | application/octet-stream              | Windows Event Viewer XML file format       |                                                                                      |
| exe                | application/x-msdownload              | Executable file for Microsoft Windows      |                                                                                      |
| iso                | application/x-iso9660-image           | ISO9660 CD/DVD image file                  |                                                                                      |
| jar                | application/java-archive              | Java archive                               |                                                                                      |
| lib                | application/octet-stream              | Windows Static Library                     | Should be compared to the linux static library format (.a) as they might be the same |
| lnk                | application/x-ms-shortcut             | Windows shell link (shortcut) file         |                                                                                      |
| msi                | application/x-msi                     | Microsoft Software Installer               |                                                                                      |
| sqlite/db/sqlitedb | application/x-sqlite3                 | SQLite Database                            |                                                                                      |
| vdi                | application/x-virtualbox-vdi          | Oracle VirtualBox Image                    |                                                                                      |
| vmdk               | application/octet-stream              | Virtual Machine Disk                       |                                                                                      |

## TODO

- .so
- .vmem Maybe (53) FF 00 F0 53 FF 00 F0 - though don't have enough sample to check
- .vmss/.vmsn
- .scr Probably not possible as ".exe" and ".scr" files are almost identical
- PKLITE / PKSFX
- .zlib
- .gz
- .xz
- Word 2.0 File: DB A5 2D 00
- What is a Perfect Office Document?? CF 11 E0 A1 B1 1A E1 00
- Microsoft OneNote note: E4 52 5C 7B 8C D8 A7 4D AE B1 53 78 D0 29 96 D3
- Windows 98 password file: E3 82 85 96
- Bitcoin Core wallet.dat file: [8 byte offset] 00 00 00 00 62 31 05 00 09 00 00 00 00 20 00 00 00 09 00 00 00 00 00 00
- Microsoft SQL Server 2000 database: 01 0F 00 00
- Windows Disk Image file: 00 00 00 00 14 00 00 00
- MultiBit Bitcoin wallet file: 0A 16 6F 72 67 2E 62 69 74 63 6F 69 6E 2E 70 72
- MultiBit Bitcoin blockchain file: 53 50 56 42
- MultiBit Bitcoin wallet information file: 6D 75 6C 74 69 42 69 74 2E 69 6E 66 6F
- Bitcoin-Qt blockchain block file: F9 BE B4 D9
