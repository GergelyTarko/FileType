package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type MIME struct {
	Type    string
	Subtype string
}

func splitMIME(s string) (string, string) {
	x := strings.Split(s, "/")
	if len(x) > 1 {
		return x[0], x[1]
	}
	return x[0], ""
}
func newMIME(mime string) MIME {
	kind, subtype := splitMIME(mime)
	return MIME{Type: kind, Subtype: subtype}
}

func (m MIME) String() string {
	return fmt.Sprintf("%s/%s", m.Type, m.Subtype)
}

type FileType struct {
	MIME        MIME
	Extension   string
	Description string
}

func (f FileType) String() string {
	return fmt.Sprintf("%s (%s) [%s]", f.Description, f.Extension, f.MIME)
}

func newType(ext string, mime string, desc string) FileType {
	return FileType{
		MIME:        newMIME(mime),
		Extension:   ext,
		Description: desc,
	}
}

func Get(path string) (FileType, error) {
	file, _ := os.OpenFile(path, os.O_RDONLY, 0666)
	defer file.Close()
	var contentByte = make([]byte, 500000)
	numByte, _ := file.Read(contentByte)
	contentByte = contentByte[:numByte]
	t, err := GetFromBuffer(contentByte)
	if err != nil && err.Error() == "Unknown filetype" {
		buf, _ := ioutil.ReadFile(path)
		t, err = GetFromBuffer(buf)
	}
	return t, err
}

// https://medium.com/asecuritysite-when-bob-met-alice/the-core-of-digital-forensics-magic-numbers-fd3e6d7a225
// https://en.wikipedia.org/wiki/List_of_file_signatures
// https://sceweb.sce.uhcl.edu/abeysekera/itec3831/labs/FILE%20SIGNATURES%20TABLE.pdf
// https://www.fileextensionguru.com/m4v
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
// https://www.wikidata.org/wiki/Q1363415
// https://www.garykessler.net/library/file_sigs.html
// https://zims-en.kiwix.campusafrica.gos.orange.com/wikipedia_en_all_nopic/A/List_of_file_signatures
func GetFromBuffer(buf []byte) (FileType, error) {

	/*  ================================================
		==================== IMAGES ====================
	    ================================================ */
	if hasPrefix(buf, 0x00, 0x00, 0x01, 0x00) { // ICO
		return newType("ico", "image/vnd.microsoft.icon", "Computer icon encoded in ICO file format"), nil

	}
	if hasPrefix(buf, 0x47, 0x49, 0x46, 0x38) { // GIF
		return newType("gif", "image/gif", "Image file encoded in the Graphics Interchange Format"), nil

	}
	if hasPrefix(buf, 0x49, 0x49, 0x2A, 0x00) || hasPrefix(buf, 0x4D, 0x4D, 0x00, 0x2A) { // TIF
		return newType("tiff", "image/tiff", "Tagged Image File Format"), nil

	}
	if hasPrefix(buf, 0xFF, 0xD8, 0xFF, 0xE0) { // JPEG/JFIF
		return newType("jpg", "image/jpeg", "JPEG using JPEG File Interchange Format (JFIF)"), nil

	}
	if hasPrefix(buf, 0xFF, 0xD8, 0xFF, 0xE1) { // JPEG/EXIF
		return newType("jpg", "image/jpeg", "JPEG using Exchangeable Image File Format (EXIF)"), nil

	}
	if hasPrefix(buf, 0xFF, 0xD8, 0xFF, 0xE8) { // JPEG/IFF
		return newType("jpg", "image/jpeg", "JPEG using Still Picture Interchange File Format"), nil

	}
	if hasPrefix(buf, 0x00, 0x00, 0x00, 0x0C, 0x6A, 0x50, 0x20, 0x20, 0x0D, 0x0A, 0x87, 0x0A) { // JPEG 2000
		if hasPrefixOffset(buf, 20, 0x6A, 0x70, 0x32) {
			return newType("jp2", "image/jp2", "JPEG 2000 format"), nil
		}
		if hasPrefixOffset(buf, 20, 0x6A, 0x70, 0x78) {
			return newType("jpx", "image/jpx", "JPEG 2000 format"), nil
		}
		if hasPrefixOffset(buf, 20, 0x6A, 0x70, 0x6D) {
			return newType("jpm", "image/jpm", "JPEG 2000 format"), nil
		}
		if hasPrefixOffset(buf, 20, 0x6D, 0x6A, 0x70, 0x32) {
			return newType("mj2", "image/mj2", "JPEG 2000 format"), nil
		}
		if hasPrefixOffset(buf, 20, 0x6A, 0x70, 0x32, 0x30) {
			return newType("jpa", "image/jpa", "JPEG 2000 format"), nil
		}
		return newType("jp2", "image/jp2", "JPEG 2000 format"), nil

	}
	if hasPrefix(buf, 0x38, 0x42, 0x50, 0x53) { // PSD
		return newType("psd", "vnd.adobe.photoshop", "Photoshop Document file, Adobe Photoshop's native file format"), nil

	}
	if hasPrefix(buf, 0x42, 0x4D) { // BMP
		return newType("bmp", "image/bmp", "A bitmap format used mostly in the Windows world"), nil

	}
	if hasPrefix(buf, 0x89, 0x50, 0x4E, 0x47) { // PNG
		return newType("png", "image/png", "Image encoded in the Portable Network Graphics format"), nil

	}

	/*  ================================================
		==================== VIDEOS ====================
	    ================================================ */
	if hasPrefix(buf, 0x1A, 0x45, 0xDF, 0xA3) && hasPrefixUndefOffset(buf, 32, []byte("matroska")) { // MATROSKA
		return newType("mkv", "video/x-matroska", "Matroska media container"), nil

	}
	if hasPrefix(buf, 0x00, 0x00, 0x1, 0xB0) { // MPEG
		return newType("mpg", "video/mpeg", "MPEG video file"), nil

	}
	if hasPrefix(buf, 0x00, 0x00, 0x01, 0xBA) { // MPEG
		return newType("mpg", "video/mpeg", "MPEG Program Stream"), nil

	}
	if hasPrefix(buf, 0x00, 0x00, 0x01, 0xB3) { // MPEG
		return newType("mpg", "video/mpeg", "MPEG-1 video and MPEG-2 video"), nil

	}
	if hasPrefixOffset(buf, 4, 0x66, 0x74, 0x79, 0x70) { // ftyp
		if hasPrefixOffset(buf, 8, []byte("qt")...) { // MOV
			return newType("mov", "video/quicktime", "ISO Media, Apple QuickTime movie, Apple QuickTime (.MOV/QT)"), nil

		}
		if hasPrefixOffset(buf, 8, []byte("M4P")...) { // M4P
			return newType("m4p", "audio/m4p", "ISO Media, Apple iTunes Audio"), nil

		}
		if hasPrefixOffset(buf, 8, []byte("M4A")...) { // M4A
			return newType("m4a", "audio/mp4", "ISO Media, Apple iTunes Audio encoded with AAC or ALAC"), nil

		}
		if hasPrefixOffset(buf, 8, []byte("M4V")...) || hasPrefixOffset(buf, 8, []byte("M4VH")...) || hasPrefixOffset(buf, 8, []byte("M4VP")...) { // M4V
			return newType("m4v", "video/x-m4v", "ISO Media, Apple iTunes Video (.M4V)"), nil

		}
		if hasPrefixOffset(buf, 8, []byte("3gp")...) { // 3GP
			return newType("3gp", "video/3gpp", "3rd Generation Partnership Project 3GPP multimedia file"), nil

		}
		// TODO There are more types to specify within ftyp
		return newType("", "", "Unknown QuickTime file format"), nil

	}
	if hasPrefixOffset(buf, 4, 0x6D, 0x6F, 0x6F, 0x76) || // MOV
		hasPrefixOffset(buf, 4, 0x66, 0x72, 0x65, 0x65) || hasPrefixOffset(buf, 4, 0x6D, 0x64, 0x61, 0x74) || hasPrefixOffset(buf, 4, 0x77, 0x69, 0x64, 0x65) ||
		hasPrefixOffset(buf, 4, 0x70, 0x6E, 0x6F, 0x74) || hasPrefixOffset(buf, 4, 0x73, 0x6B, 0x69, 0x70) {
		return newType("mov", "video/quicktime", "ISO Media, Apple QuickTime movie, Apple QuickTime (.MOV/QT)"), nil

	}
	if hasPrefix(buf, 0x1A, 0x45, 0xDF, 0xA3) && hasPrefixUndefOffset(buf, 32, []byte("webm")) { // WEBM
		return newType("webm", "video/webm", "Matroska media container (webm)"), nil

	}
	if hasPrefix(buf, 0x30, 0x26, 0xB2, 0x75, 0x8E, 0x66, 0xCF, 0x11, 0xA6, 0xD9, 0x00, 0xAA, 0x00, 0x62, 0xCE, 0x6C) { // WMV
		return newType("wmv", "video/x-ms-wmv", "Microsoft Windows Media Audio/Video File"), nil

	}
	if hasPrefix(buf, 0x46, 0x4C, 0x56, 0x01) { // FLV
		return newType("flv", "video/x-flv", "Flash Video file"), nil

	}

	/*  ================================================
		================ AUDIO + IFF + RIFF ============
	    ================================================ */
	if hasPrefix(buf, 0x46, 0x4F, 0x52, 0x4D) { // IFF
		if hasPrefixOffset(buf, 8, 0x41, 0x49, 0x46, 0x46) { // AIFF
			return newType("aiff", "audio/x-aiff", "Audio Interchange File Format"), nil
		}
		return newType("", "", "Unknown IFF format"), nil
		// TODO Implement the rest of the IFF formats
	}
	if hasPrefix(buf, 0x52, 0x49, 0x46, 0x46) { // RIFF
		if hasPrefixOffset(buf, 8, 0x57, 0x41, 0x56, 0x45) { // WAV
			return newType("wav", "audio/x-wav", "Waveform Audio File Format"), nil

		}
		if hasPrefixOffset(buf, 8, 0x41, 0x56, 0x49, 0x20) { // AVI
			return newType("avi", "video/x-msvideo", "Audio Video Interleave video format"), nil

		}
		if hasPrefixOffset(buf, 8, 0x57, 0x45, 0x42, 0x50) { // WEBP
			return newType("webp", "image/webp", "Google WebP image file"), nil

		}
		// TODO Implement the rest of the RIFF formats

	}
	if hasPrefix(buf, 0x66, 0x4C, 0x61, 0x43) { // FLAC
		return newType("flac", "audio/x-flac", "Free Lossless Audio Codec"), nil
	}
	if hasPrefix(buf, 0x4F, 0x67, 0x67, 0x53) { // OGG
		if hasPrefixOffset(buf, 28, 0x7F, 0x46, 0x4C, 0x41, 0x43) { // flac
			return newType("oga", "audio/ogg", "An open source media container format (Audio)"), nil
		}
		if hasPrefixOffset(buf, 28, 0x01, 0x76, 0x6F, 0x72, 0x62, 0x69, 0x73) { // vorbis
			return newType("ogg", "audio/ogg", "An open source media container format (Audio)"), nil
		}
		if hasPrefixOffset(buf, 28, 0x01, 0x76, 0x69, 0x64, 0x65, 0x6F, 0x00) { // video
			return newType("ogm", "video/ogg", "OGG Media Stream file"), nil
		}
		if hasPrefixOffset(buf, 28, 0x80, 0x74, 0x68, 0x65, 0x6F, 0x72, 0x61) { // theora
			return newType("ogv", "video/ogg", "Ogg Video file"), nil
		}
		if hasPrefixOffset(buf, 28, 0x4F, 0x70, 0x75, 0x73, 0x48, 0x65, 0x61, 0x64) { // opus
			return newType("opus", "audio/opus", "An open source media container format (Audio)"), nil
		}
		if hasPrefixOffset(buf, 28, 0x53, 0x70, 0x65, 0x65, 0x78, 0x20, 0x20) { // speex
			return newType("spx", "audio/ogg", "An open source media container format (Audio)"), nil
		}
		return newType("ogx", "application/ogg", "An open source media container format"), nil
	}
	if hasPrefix(buf, 0x4D, 0x54, 0x68, 0x64) { // MIDI
		return newType("mid", "audio/midi", "MIDI sound file"), nil
	}

	/*  ================================================
		==================== ARCHIVES ==================
	    ================================================ */
	if hasPrefix(buf, 0x1F, 0x9D) { // Z/LZW
		return newType("Z", "application/x-compress", "Compressed tape archive file using standard Lempel-Ziv-Welch compression"), nil

	}
	if hasPrefix(buf, 0x1F, 0xA0) { //Z/LZH
		return newType("Z", "application/x-compress", "Compressed tape archive file using Lempel-Ziv-Huffman compression"), nil
	}
	if hasPrefix(buf, 0x50, 0x4B, 0x03, 0x04) { // PK
		/*
			zip  	+
			aar
			apk		-
			docx	+
			epub 	+
			ipa
			jar		+
			kmz
			maff
			odp		+
			ods		+
			odt		+
			pk3
			pk4
			pptx	+
			usdz
			vsdx
			xlsx	+
			xpi		+	*/
		fileNameLength := binary.LittleEndian.Uint16(buf[26:28])
		fileName := string(buf[30 : 30+(fileNameLength)])
		if fileName == "mimetype" {
			if hasPrefixOffset(buf, 38, []byte("application/epub+zip")...) { // EPUB
				return newType("epub", "application/epub+zip", "E-book file format"), nil
			}
			if hasPrefixOffset(buf, 38, []byte("application/vnd.oasis.opendocument.text")...) { // ODT
				return newType("odt", "application/vnd.oasis.opendocument.text", "OpenOffice Text Document"), nil
			}
			if hasPrefixOffset(buf, 38, []byte("application/vnd.oasis.opendocument.presentation")...) { // ODP
				return newType("odp", "application/vnd.oasis.opendocument.presentation", "OpenOffice Presentation Document"), nil
			}
			if hasPrefixOffset(buf, 38, []byte("application/vnd.oasis.opendocument.spreadsheet")...) { // ODS
				return newType("ods", "application/vnd.oasis.opendocument.spreadsheet", "OpenDocument Spreadsheet Document"), nil
			}
		} else if strings.HasSuffix(fileName, ".xml") || strings.HasSuffix(fileName, ".rels") {
			if hasPrefixUndefOffset(buf, len(buf)-5, []byte("word/")) { // MS-WORD
				return newType("docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "Office Open XML File Format"), nil
			}
			if hasPrefixUndefOffset(buf, len(buf)-5, []byte("ppt/")) { // MS-POWERPOINT
				return newType("pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation", "Office Open XML File Format"), nil
			}
			if hasPrefixUndefOffset(buf, len(buf)-5, []byte("xl/")) { // MS-EXCEL
				return newType("xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "Office Open XML File Format"), nil
			}
		} else if fileName == "META-INF/mozilla.rsa" { // MOZILLA PLUGIN
			return newType("xpi", "application/x-xpinstall", "Mozilla Browser Archive"), nil
		} else if hasPrefixUndefOffset(buf, 512, []byte("META-INF/MANIFEST.MF")) { // TODO APK OR JAR
			if hasPrefixUndefOffset(buf, len(buf)-len([]byte("AndroidManifest.xml")), []byte("AndroidManifest.xml")) { // APK 1
				return newType("apk", "application/apk-archive", "Android package"), nil
			}
			return newType("jar", "application/java-archive", "Java archive"), nil // Default to JAR but it can be apk as well
		} else if 1 == 0 { // TODO APK
			return newType("apk", "application/apk-archive", "Android package"), nil
		}
		return newType("zip", "application/zip", "PKZIP archive file"), nil
	}
	if hasPrefix(buf, 0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x00) { // RAR v4.x
		return newType("rar", "application/vnd.rar", "Roshal ARchive compressed archive v1.50 onwards"), nil
	}
	if hasPrefix(buf, 0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x01, 0x00) { // RAR v5.x
		return newType("rar", "application/vnd.rar", "Roshal ARchive compressed archive v5.00 onwards"), nil
	}
	if hasPrefixOffset(buf, 257, 0x75, 0x73, 0x74, 0x61, 0x72, 0x00, 0x30, 0x30) || hasPrefixOffset(buf, 257, 0x75, 0x73, 0x74, 0x61, 0x72, 0x20, 0x20, 0x00) { // TAR
		return newType("tar", "application/x-tar", "tar archive"), nil
	}
	if hasPrefix(buf, 0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C) { // 7Z
		return newType("7z", "application/x-7z-compressed", "7-Zip File Format"), nil
	}
	if hasPrefix(buf, 0x21, 0x3C, 0x61, 0x72, 0x63, 0x68, 0x3E, 0x0A) { // DEB
		return newType("deb", "application/vnd.debian.binary-package", "Linux deb file"), nil
	}
	if hasPrefixOffset(buf, 32769, 0x43, 0x44, 0x30, 0x30, 0x31) || hasPrefixOffset(buf, 34817, 0x43, 0x44, 0x30, 0x30, 0x31) || hasPrefixOffset(buf, 36865, 0x43, 0x44, 0x30, 0x30, 0x31) { //ISO
		return newType("iso", "application/x-iso9660-image", "ISO9660 CD/DVD image file"), nil
	}
	/*  ================================================
		==================== DOCUMENTS =================
	    ================================================ */
	if hasPrefix(buf, 0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1) { //CFB
		if hasPrefixOffset(buf, 512, 0xEC, 0xA5, 0xC1, 0x00) { // DOC
			return newType("doc", "application/msword", "MS Office Word document file"), nil
		}
		if hasPrefixOffset(buf, 512, 0x00, 0x6E, 0x1E, 0xF0) || hasPrefixOffset(buf, 512, 0xA0, 0x46, 0x1D, 0xF0) || hasPrefixOffset(buf, 512, 0x0F, 0x00, 0xE8, 0x03) || hasPrefixUndefOffset(buf, len(buf)-38, []byte{0x50, 0x00, 0x6F, 0x00, 0x77, 0x00, 0x65, 0x00, 0x72, 0x00, 0x50, 0x00, 0x6F, 0x00, 0x69, 0x00, 0x6E, 0x00, 0x74, 0x00, 0x20, 0x00, 0x44, 0x00, 0x6F, 0x00, 0x63, 0x00, 0x75, 0x00, 0x6D, 0x00, 0x65, 0x00, 0x6E, 0x00, 0x74, 0x00}) { // PPT
			return newType("ppt", "application/vnd.ms-powerpoint", "MS Office PowerPoint presentation file"), nil
		}
		if hasPrefixOffset(buf, 512, 0x09, 0x08, 0x10, 0x00, 0x00, 0x06, 0x05, 0x00) { // XLS
			return newType("xls", "application/vnd.ms-excel", "MS Office Word document file"), nil
		}
		if hasPrefix(buf, 0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0x00, 0x04, 0x00, 0xFE, 0xFF, 0x0C, 0x00, 0x06) || // 64-bit msi?
			hasPrefix(buf, 0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0x00, 0x03, 0x00, 0xFE, 0xFF, 0x09, 0x00, 0x06) { // 32-bit msi?
			return newType("msi", "application/x-msi", "Microsoft Software Installer"), nil
		}
	}
	/*  ================================================
		==================== Other======================
	    ================================================ */
	if hasPrefix(buf, 0xCA, 0xFE, 0xBA, 0xBE) { // JAVA CLASS
		return newType("class", "application/x-java-class", "Java class file"), nil
	}
	if hasPrefix(buf, 0x4D, 0x5A) && len(buf) > 64 { // PE FILE http://www.pelib.com/resources/luevel.txt
		e_lfanew := binary.LittleEndian.Uint32(buf[60:64]) // PE Header location
		if len(buf) > int(e_lfanew) {
			characteristic := binary.LittleEndian.Uint16(buf[e_lfanew+22 : e_lfanew+24])
			if characteristic&0x2000 > 0 { // DLL
				return newType("dll", "application/x-msdownload", "Dynamic-link library for Microsoft Windows"), nil
			}
			if characteristic&0x0002 > 0 { // Is valid executable
				return newType("exe", "application/x-msdownload", "Executable file for Microsoft Windows"), nil
			}
			// TODO Check whether an executable image is 32-bit or 64-bit using the optional header: https://learn.microsoft.com/en-us/windows/win32/debug/pe-format#optional-header-image-only
			// https://tech-zealots.com/malware-analysis/pe-portable-executable-structure-malware-analysis-part-2/
		}
		//return newType("exe", "application/x-msdownload", "Executable file for Microsoft Windows"), nil
	}
	if hasPrefix(buf, 0x25, 0x50, 0x44, 0x46) { // PDF
		return newType("pdf", "application/pdf", "Portable Document Format"), nil
	}
	if hasPrefix(buf, 0x3C, 0x3C, 0x3C, 0x20, 0x4F, 0x72, 0x61, 0x63, 0x6C, 0x65, 0x20, 0x56, 0x4D, 0x20, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6C, 0x42, 0x6F, 0x78, 0x20, 0x44, 0x69, 0x73, 0x6B, 0x20, 0x49, 0x6D, 0x61, 0x67, 0x65, 0x20, 0x3E, 0x3E, 0x3E) { //VDI
		// Note: The value tends to change every time the VirtualBox product name or owner changes, or if the VDI was created by a third party app. (source: https://forums.virtualbox.org/viewtopic.php?f=1&t=47732)
		return newType("vdi", "application/x-virtualbox-vdi", "Oracle VirtualBox Disk Image"), nil
	}
	if hasPrefix(buf, 0x4C, 0x00, 0x00, 0x00, 0x01, 0x14, 0x02, 0x00) { // LNK
		return newType("lnk", "application/x-ms-shortcut", "Windows shell link (shortcut) file"), nil
	}
	return FileType{}, errors.New("Unknown filetype")
}

func hasPrefixOffset(buf []byte, offset int, prefix ...byte) bool {
	if len(buf) < len(prefix)+offset {
		return false
	}
	for i, b := range prefix {
		if buf[offset+i] != b {
			return false
		}
	}
	return true
}

func hasPrefix(buf []byte, prefix ...byte) bool {
	if len(buf) < len(prefix) {
		return false
	}
	for i, b := range prefix {
		if buf[i] != b {
			return false
		}
	}
	return true
}

func hasPrefixUndefOffset(buf []byte, maxOffset int, seq []byte) bool {
	if len(buf) < len(seq)+maxOffset {
		return false
	}

	index := bytes.Index(buf[:len(seq)+maxOffset], seq)
	if index == -1 {
		return false
	}
	return true
}
