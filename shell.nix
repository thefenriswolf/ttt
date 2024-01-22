{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.gotools # godoc, ...
    pkgs.go-tools # staticcheck, ...
    pkgs.delve
    pkgs.gopls
    pkgs.gomodifytags
    pkgs.gore
    pkgs.gotests
    pkgs.gocode
    pkgs.govulncheck
    # keep this line if you use bash
    pkgs.bashInteractive
  ];
}
