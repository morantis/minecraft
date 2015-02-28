// textures_test.go
package minecraft

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTextures(t *testing.T) {

	Convey("Test decodeTextureProperty", t, func() {

		Convey("Should correctly decode Skin and Cape URL", func() {
			// citricsquid
			sessionProfileProperty := SessionProfileProperty{Name: "textures", Value: "eyJ0aW1lc3RhbXAiOjE0MjQ5ODM2MTI1NzgsInByb2ZpbGVJZCI6IjQ4YTBhN2U0ZDU1OTQ4NzNhNjE3ZGMxODlmNzZhOGExIiwicHJvZmlsZU5hbWUiOiJjaXRyaWNzcXVpZCIsInRleHR1cmVzIjp7IlNLSU4iOnsidXJsIjoiaHR0cDovL3RleHR1cmVzLm1pbmVjcmFmdC5uZXQvdGV4dHVyZS9lMWM2YzliNmRlODhmNDE4OGY5NzMyOTA5Yzc2ZGZjZDdiMTZhNDBhMDMxY2UxYjQ4NjhlNGQxZjg4OThlNGYifSwiQ0FQRSI6eyJ1cmwiOiJodHRwOi8vdGV4dHVyZXMubWluZWNyYWZ0Lm5ldC90ZXh0dXJlL2MzYWY3ZmI4MjEyNTQ2NjQ1NThmMjgzNjExNThjYTczMzAzYzlhODVlOTZlNTI1MTEwMjk1OGQ3ZWQ2MGM0YTMifX19=="}
			sessionProfile := SessionProfileResponse{Properties: []SessionProfileProperty{sessionProfileProperty}}

			profileTextureProperty, err := decodeTextureProperty(sessionProfile)

			So(err, ShouldBeNil)
			So(profileTextureProperty.Textures.Skin.URL, ShouldEqual, "http://textures.minecraft.net/texture/e1c6c9b6de88f4188f9732909c76dfcd7b16a40a031ce1b4868e4d1f8898e4f")
			So(profileTextureProperty.Textures.Cape.URL, ShouldEqual, "http://textures.minecraft.net/texture/c3af7fb821254664558f28361158ca73303c9a85e96e5251102958d7ed60c4a3")
		})

		Convey("Should only decode Skin URL", func() {
			// citricsquid
			sessionProfileProperty := SessionProfileProperty{Name: "textures", Value: "eyJ0aW1lc3RhbXAiOjE0MjQ5ODM2MTI1NzgsInByb2ZpbGVJZCI6IjQ4YTBhN2U0ZDU1OTQ4NzNhNjE3ZGMxODlmNzZhOGExIiwicHJvZmlsZU5hbWUiOiJjaXRyaWNzcXVpZCIsInRleHR1cmVzIjp7IlNLSU4iOnsidXJsIjoiaHR0cDovL3RleHR1cmVzLm1pbmVjcmFmdC5uZXQvdGV4dHVyZS9lMWM2YzliNmRlODhmNDE4OGY5NzMyOTA5Yzc2ZGZjZDdiMTZhNDBhMDMxY2UxYjQ4NjhlNGQxZjg4OThlNGYifX19"}
			sessionProfile := SessionProfileResponse{Properties: []SessionProfileProperty{sessionProfileProperty}}

			profileTextureProperty, err := decodeTextureProperty(sessionProfile)

			So(err, ShouldBeNil)
			So(profileTextureProperty.Textures.Skin.URL, ShouldEqual, "http://textures.minecraft.net/texture/e1c6c9b6de88f4188f9732909c76dfcd7b16a40a031ce1b4868e4d1f8898e4f")
			So(profileTextureProperty.Textures.Cape.URL, ShouldBeBlank)
		})

		Convey("Should only decode Cape URL", func() {
			// citricsquid
			sessionProfileProperty := SessionProfileProperty{Name: "textures", Value: "eyJ0aW1lc3RhbXAiOjE0MjQ5ODM2MTI1NzgsInByb2ZpbGVJZCI6IjQ4YTBhN2U0ZDU1OTQ4NzNhNjE3ZGMxODlmNzZhOGExIiwicHJvZmlsZU5hbWUiOiJjaXRyaWNzcXVpZCIsInRleHR1cmVzIjp7IkNBUEUiOnsidXJsIjoiaHR0cDovL3RleHR1cmVzLm1pbmVjcmFmdC5uZXQvdGV4dHVyZS9jM2FmN2ZiODIxMjU0NjY0NTU4ZjI4MzYxMTU4Y2E3MzMwM2M5YTg1ZTk2ZTUyNTExMDI5NThkN2VkNjBjNGEzIn19fQ=="}
			sessionProfile := SessionProfileResponse{Properties: []SessionProfileProperty{sessionProfileProperty}}

			profileTextureProperty, err := decodeTextureProperty(sessionProfile)

			So(err, ShouldBeNil)
			So(profileTextureProperty.Textures.Skin.URL, ShouldBeBlank)
			So(profileTextureProperty.Textures.Cape.URL, ShouldEqual, "http://textures.minecraft.net/texture/c3af7fb821254664558f28361158ca73303c9a85e96e5251102958d7ed60c4a3")
		})

		Convey("Should error about no textures", func() {
			sessionProfile := SessionProfileResponse{}

			profileTextureProperty, err := decodeTextureProperty(sessionProfile)

			So(err.Error(), ShouldContainSubstring, "No textures property.")
			So(profileTextureProperty, ShouldResemble, SessionProfileTextureProperty{})
		})

		Convey("Should error trying to decode", func() {
			sessionProfileProperty := SessionProfileProperty{Name: "textures", Value: ""}
			sessionProfile := SessionProfileResponse{Properties: []SessionProfileProperty{sessionProfileProperty}}

			profileTextureProperty, err := decodeTextureProperty(sessionProfile)

			So(err.Error(), ShouldContainSubstring, "Error decoding texture property")
			So(profileTextureProperty, ShouldResemble, SessionProfileTextureProperty{})
		})

	})

	Convey("Test Texture.fetch", t, func() {

		Convey("clone1018 texture should return the correct skin", func() {
			texture := &Texture{URL: "http://textures.minecraft.net/texture/cd9ca55e9862f003ebfa1872a9244ad5f721d6b9e6883dd1d42f87dae127649"}

			err := texture.fetch()

			So(err, ShouldBeNil)
			So(texture.Hash, ShouldEqual, "a04a26d10218668a632e419ab073cf57")
		})

		Convey("Bad texture requests should gracefully fail", func() {

			Convey("No texture URL", func() {
				texture := &Texture{URL: ""}

				err := texture.fetch()

				So(err.Error(), ShouldContainSubstring, "No Texture URL")
			})

			Convey("Bad texture URL (non-200)", func() {
				texture := &Texture{URL: "http://textures.minecraft.net/texture/"}

				err := texture.fetch()

				So(err.Error(), ShouldContainSubstring, "Error retrieving texture")
			})

			Convey("Bad texture URL (non-image)", func() {
				texture := &Texture{URL: "http://google.com"}

				err := texture.fetch()

				So(err.Error(), ShouldContainSubstring, "image: unknown format")
			})

			Convey("Not a URL", func() {
				texture := &Texture{URL: "//"}

				err := texture.fetch()

				So(err.Error(), ShouldContainSubstring, "Unable to Get URL")
			})

		})

	})

	Convey("Test Steve", t, func() {

		Convey("Steve should return valid image", func() {
			steveImg, err := FetchImageForSteve()

			So(err, ShouldBeNil)
			So(steveImg, ShouldNotBeNil)
		})

		Convey("Steve should return valid skin", func() {
			steveSkin, err := FetchSkinForSteve()

			So(err, ShouldBeNil)
			So(steveSkin, ShouldNotResemble, &Skin{})
			So(steveSkin.Hash, ShouldEqual, "98903c1609352e11552dca79eb1ce3d6")
		})

	})

	Convey("Test Skins", t, func() {

		// Must be careful to not request same profile from session server more than once per ~30 seconds
		Convey("d9135e082f2244c89cb0bee234155292 should return valid image from Mojang", func() {
			// clone1018
			skin, err := FetchSkinUUID("d9135e082f2244c89cb0bee234155292")

			So(err, ShouldBeNil)
			So(skin, ShouldNotResemble, &Skin{})
			So(skin.Hash, ShouldEqual, "a04a26d10218668a632e419ab073cf57")
		})

		// Must be careful to not request same profile from session server more than once per ~30 seconds
		Convey("00000000000000000000000000000000 err from Mojang", func() {
			// clone1018
			skin, err := FetchSkinUUID("00000000000000000000000000000000")

			So(err.Error(), ShouldContainSubstring, "User not found")
			So(skin, ShouldResemble, &Skin{})
		})

		Convey("clone1018 should return valid image from Mojang", func() {
			skin, err := FetchSkinUsernameMojang("clone1018")

			So(err, ShouldBeNil)
			So(skin, ShouldNotResemble, &Skin{})
			So(skin.Hash, ShouldEqual, "a04a26d10218668a632e419ab073cf57")
		})

		Convey("Wooxye should err from Mojang", func() {
			skin, err := FetchSkinUsernameMojang("Wooxye")

			So(err.Error(), ShouldContainSubstring, "Texture not found")
			So(skin, ShouldResemble, &Skin{Texture{Source: "Mojang", URL: "http://skins.minecraft.net/MinecraftSkins/Wooxye.png"}})
		})

		Convey("clone1018 should return valid image from S3", func() {
			skin, err := FetchSkinUsernameS3("clone1018")

			So(err, ShouldBeNil)
			So(skin, ShouldNotResemble, &Skin{})
			So(skin.Hash, ShouldEqual, "a04a26d10218668a632e419ab073cf57")
		})

		Convey("Wooxye should err from S3", func() {
			skin, err := FetchSkinUsernameS3("Wooxye")

			So(err.Error(), ShouldContainSubstring, "Texture not found")
			So(skin, ShouldResemble, &Skin{Texture{Source: "S3", URL: "http://s3.amazonaws.com/MinecraftSkins/Wooxye.png"}})
		})

	})

	// Must be careful to not request same profile from session server more than once per ~30 seconds
	Convey("Test Capes", t, func() {

		Convey("48a0a7e4d5594873a617dc189f76a8a1 should return a Cape from Mojang", func() {
			// citricsquid
			cape, err := FetchCapeUUID("48a0a7e4d5594873a617dc189f76a8a1")

			So(err, ShouldBeNil)
			So(cape, ShouldNotResemble, &Cape{Texture{Source: "SessionProfile"}})
			So(cape.Hash, ShouldEqual, "8cbf8786caba2f05383cf887be592ee6")
		})

		Convey("2f3665cc5e29439bbd14cb6d3a6313a7 should err from Mojang (No Cape)", func() {
			// lukegb
			cape, err := FetchCapeUUID("2f3665cc5e29439bbd14cb6d3a6313a7")

			So(err.Error(), ShouldContainSubstring, "Cape URL is not present.")
			So(cape, ShouldResemble, &Cape{Texture{Source: "SessionProfile"}})
			So(cape.Hash, ShouldBeBlank)
		})

		Convey("00000000000000000000000000000001 should err from Mojang (No User)", func() {
			cape, err := FetchCapeUUID("00000000000000000000000000000001")

			So(err.Error(), ShouldContainSubstring, "User not found")
			So(cape, ShouldResemble, &Cape{})
		})

		Convey("citricsquid should return a Cape from Mojang", func() {
			cape, err := FetchCapeUsernameMojang("citricsquid")

			So(err, ShouldBeNil)
			So(cape, ShouldNotResemble, &Cape{})
			So(cape.Hash, ShouldEqual, "8cbf8786caba2f05383cf887be592ee6")
		})

		Convey("Wooxye should err from Mojang", func() {
			cape, err := FetchCapeUsernameMojang("Wooxye")

			So(err.Error(), ShouldContainSubstring, "Texture not found")
			So(cape, ShouldResemble, &Cape{Texture{Source: "Mojang", URL: "http://skins.minecraft.net/MinecraftCloaks/Wooxye.png"}})
		})

		Convey("citricsquid should return a Cape from S3", func() {
			cape, err := FetchCapeUsernameS3("citricsquid")

			So(err, ShouldBeNil)
			So(cape, ShouldNotResemble, &Cape{})
			So(cape.Hash, ShouldEqual, "8cbf8786caba2f05383cf887be592ee6")
		})

		Convey("Wooxye should err from S3", func() {
			cape, err := FetchCapeUsernameS3("Wooxye")

			So(err.Error(), ShouldContainSubstring, "Texture not found")
			So(cape, ShouldResemble, &Cape{Texture{Source: "S3", URL: "http://s3.amazonaws.com/MinecraftCloaks/Wooxye.png"}})
		})

	})

}
