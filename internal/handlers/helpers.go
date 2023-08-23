package handlers

import (
	"github.com/diazharizky/teleforecaster/pkg/airvisual"
)

func resolveErrMessage(err error) (msg string) {
	switch err {
	case airvisual.StateNotSupportedError:
		msg = "Provinsi tidak didukung."
	case airvisual.CityNotSupportedError:
		msg = "Kota/kabupaten tidak didukung."
	case airvisual.RateLimitError:
		msg = "Gagal mendapatkan data, silahkan coba lagi."
	default:
		msg = "Terjadi kesalahan, silahkan coba lagi."
	}

	return
}
