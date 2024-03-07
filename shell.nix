{ pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.gotools # godoc, ...
    pkgs.go-tools # staticcheck, ...
    pkgs.delve
    pkgs.gopls
    pkgs.gcc
    pkgs.gomodifytags
    pkgs.gore
    pkgs.gotests
    pkgs.gocode
    pkgs.govulncheck
    pkgs.revive
    # pkgs.goreleaser

    pkgs.gcc
    pkgs.bashInteractive
  ];
}
