class GofiberCreator < Formula
  desc "CLI tool to scaffold production-ready Go Fiber projects"
  homepage "https://github.com/songqii/go_base_temp"
  version "1.0.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/songqii/go_base_temp/releases/download/v#{version}/gofiber-creator-darwin-arm64"
      sha256 "PLACEHOLDER_ARM64"

      def install
        bin.install "gofiber-creator-darwin-arm64" => "gofiber-creator"
      end
    else
      url "https://github.com/songqii/go_base_temp/releases/download/v#{version}/gofiber-creator-darwin-amd64"
      sha256 "PLACEHOLDER_AMD64"

      def install
        bin.install "gofiber-creator-darwin-amd64" => "gofiber-creator"
      end
    end
  end

  on_linux do
    url "https://github.com/songqii/go_base_temp/releases/download/v#{version}/gofiber-creator-linux-amd64"
    sha256 "PLACEHOLDER_LINUX"

    def install
      bin.install "gofiber-creator-linux-amd64" => "gofiber-creator"
    end
  end

  test do
    assert_match "gofiber-creator version", shell_output("#{bin}/gofiber-creator -v")
  end
end
