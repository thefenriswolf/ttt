{ pkgs ? (let sources = import ./nix/sources.nix;
in import sources.nixpkgs {
  overlays = [ (import "${sources.gomod2nix}/overlay.nix") ];
}) }:

pkgs.buildGoApplication {
  pname = "ttt";
  version = "20240119";
  pwd = ./.;
  src = ./.;
  modules = ./gomod2nix.toml;

  meta = {
    description = "Time Tracker Tool written in Go";
    homepage = "https://github.com/thefenriswolf/ttt";
    license = pkgs.lib.licenses.bsd3;
    maintainers = with pkgs.lib.maintainers; [ thefenriswolf ];
  };
}
