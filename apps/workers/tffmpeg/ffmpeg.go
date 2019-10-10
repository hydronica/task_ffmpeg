package main

type MediaInfo struct {
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                    string `json:"@type"`
			VideoCount              string `json:"VideoCount,omitempty"`
			AudioCount              string `json:"AudioCount,omitempty"`
			FileExtension           string `json:"FileExtension,omitempty"`
			Format                  string `json:"Format"`
			FormatProfile           string `json:"Format_Profile,omitempty"`
			CodecID                 string `json:"CodecID"`
			CodecIDCompatible       string `json:"CodecID_Compatible,omitempty"`
			FileSize                string `json:"FileSize,omitempty"`
			Duration                string `json:"Duration"`
			OverallBitRate          string `json:"OverallBitRate,omitempty"`
			FrameRate               string `json:"FrameRate,omitempty"`
			FrameCount              string `json:"FrameCount"`
			StreamSize              string `json:"StreamSize"`
			HeaderSize              string `json:"HeaderSize,omitempty"`
			DataSize                string `json:"DataSize,omitempty"`
			FooterSize              string `json:"FooterSize,omitempty"`
			IsStreamable            string `json:"IsStreamable,omitempty"`
			Title                   string `json:"Title,omitempty"`
			Movie                   string `json:"Movie,omitempty"`
			EncodedDate             string `json:"Encoded_Date"`
			TaggedDate              string `json:"Tagged_Date"`
			FileModifiedDate        string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal   string `json:"File_Modified_Date_Local,omitempty"`
			EncodedApplication      string `json:"Encoded_Application,omitempty"`
			StreamOrder             string `json:"StreamOrder,omitempty"`
			ID                      string `json:"ID,omitempty"`
			FormatLevel             string `json:"Format_Level,omitempty"`
			FormatSettingsCABAC     string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsRefFrames string `json:"Format_Settings_RefFrames,omitempty"`
			BitRate                 string `json:"BitRate,omitempty"`
			Width                   string `json:"Width,omitempty"`
			Height                  string `json:"Height,omitempty"`
			StoredHeight            string `json:"Stored_Height,omitempty"`
			SampledWidth            string `json:"Sampled_Width,omitempty"`
			SampledHeight           string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio        string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio      string `json:"DisplayAspectRatio,omitempty"`
			Rotation                string `json:"Rotation,omitempty"`
			FrameRateMode           string `json:"FrameRate_Mode,omitempty"`
			FrameRateModeOriginal   string `json:"FrameRate_Mode_Original,omitempty"`
			ChromaSubsampling       string `json:"ChromaSubsampling,omitempty"`
			BitDepth                string `json:"BitDepth,omitempty"`
			ScanType                string `json:"ScanType,omitempty"`
			EncodedLibrary          string `json:"Encoded_Library,omitempty"`
			EncodedLibraryName      string `json:"Encoded_Library_Name,omitempty"`
			EncodedLibraryVersion   string `json:"Encoded_Library_Version,omitempty"`
			EncodedLibrarySettings  string `json:"Encoded_Library_Settings,omitempty"`
			Extra                   struct {
				CodecConfigurationBox string `json:"Codec_configuration_box"`
			} `json:"extra,omitempty"`
			BitRateMode          string `json:"BitRate_Mode,omitempty"`
			SamplingRate         string `json:"SamplingRate,omitempty"`
			SamplingCount        string `json:"SamplingCount,omitempty"`
			StreamSizeProportion string `json:"StreamSize_Proportion,omitempty"`
			Default              string `json:"Default,omitempty"`
			AlternateGroup       string `json:"AlternateGroup,omitempty"`
		} `json:"track"`
	} `json:"media"`
}
