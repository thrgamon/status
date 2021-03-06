# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Status < Formula
  desc "A simple tool for getting information on http status codes."
  homepage "https://github.com/thrgamon/status"
  version "0.1.1"
  bottle :unneeded

  if OS.mac? && Hardware::CPU.intel?
    url "https://github.com/thrgamon/status/releases/download/v0.1.1/status_0.1.1_Darwin_x86_64.tar.gz"
    sha256 "0c65b075c191b3805c851c28a9cb65a820e545351b2b54c51770f1577b6728f4"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/thrgamon/status/releases/download/v0.1.1/status_0.1.1_Linux_x86_64.tar.gz"
    sha256 "dddf24f505d227564a4586be61a1a77f5471031d94f2b8afdfc0869e12b8873a"
  end
  if OS.linux? && Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
    url "https://github.com/thrgamon/status/releases/download/v0.1.1/status_0.1.1_Linux_arm64.tar.gz"
    sha256 "a40bc75b61abbe2dcb4ed33873105a17b8597b7c3c325e942dbe1c965f961a3e"
  end

  def install
    bin.install "status"
  end
end
