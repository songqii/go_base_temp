class GofiberCreator < Formula
  desc "CLI tool to scaffold production-ready Go Fiber projects"
  homepage "https://github.com/songqii/go_base_temp"
  version "1.0.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/songqii/go_base_temp/releases/download/v#{version}/gofiber-creator-darwin-arm64"
      sha256 "bb7e02096c86c58fadfb9f8c9ebe0eb23313c4581f1aa722c4db26249745e3de"

      def install
        bin.install "gofiber-creator-darwin-arm64" => "gofiber-creator"
      end
    else
      url "https://github.com/songqii/go_base_temp/releases/download/v#{version}/gofiber-creator-darwin-amd64"
      sha256 "71ea9935eaaee25a3f7269e40feee51ee37cd2f7048bbe485d0831f4d348b678"

      def install
        bin.install "gofiber-creator-darwin-amd64" => "gofiber-creator"
      end
    end
  end

  on_linux do
    url "https://github.com/songqii/go_base_temp/releases/download/v#{version}/gofiber-creator-linux-amd64"
    sha256 "e397ee974638b5aad745edcc516bbb3b5c865394eb920ba4b465df2d0fa98ab0"

    def install
      bin.install "gofiber-creator-linux-amd64" => "gofiber-creator"
    end
  end

  test do
    assert_match "gofiber-creator version", shell_output("#{bin}/gofiber-creator -v")
  end
end
