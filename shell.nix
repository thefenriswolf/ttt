{ pkgs ? (let sources = import ./nix/sources.nix;
in import sources.nixpkgs {
  overlays = [ (import "${sources.gomod2nix}/overlay.nix") ];
}) }:

let goEnv = pkgs.mkGoEnv { pwd = ./.; };
in pkgs.mkShell {
  packages = [
    goEnv
    pkgs.gomod2nix
    pkgs.niv
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
  ];
}
