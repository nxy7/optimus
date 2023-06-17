{
  description = "Project starter";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flakeUtils.url = "github:numtide/flake-utils";
    nix2container.url = "github:nlewo/nix2container";
  };

  outputs = { self, nixpkgs, flakeUtils, nix2container, ... }@inputs:
    flakeUtils.lib.eachSystem [ "x86_64-linux" ] (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        };
      in {
        devShell = pkgs.mkShell { packages = with pkgs; [ go cobra-cli ]; };
        defaultPackage = pkgs.buildGoModule rec {
          pname = "optimus";
          version = "0.0.1";

          src = ./.;

          vendorHash = "sha256-JFvC9V0xS8SZSdLsOtpyTrFzXjYAOaPQaJHdcnJzK3s=";

          meta = with pkgs.lib; {
            description = "Simple command-line snippet manager, written in Go";
            homepage = "https://github.com/nxy7/optimus";
            license = licenses.mit;
            maintainers = with maintainers; [ nxyt ];
          };
        };
      });
}
