package media

func (f Format) String() (str string) {
    switch f {
    case Format_MP3:
        str = `mp3`
    case Format_MP4:
        str = `mp4`
    case Format_FORMAT_UNSPECIFIED:
        str = `unknown`
    default:
        str = `error`
    }

    return
}
