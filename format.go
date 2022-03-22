package media

func (f Format) Name() (name string) {
    switch f {
    case Format_MP3:
        name = `mp3`
    case Format_MP4:
        name = `mp4`
    case Format_FORMAT_UNSPECIFIED:
        name = `unknown`
    default:
        name = `error`
    }

    return
}
