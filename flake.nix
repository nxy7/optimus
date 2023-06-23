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
        devShell =
          pkgs.mkShell { packages = with pkgs; [ go gopls cobra-cli ]; };
        defaultPackage = pkgs.buildGoModule rec {
          pname = "optimus";
          version = "0.0.1";

          # we're looking for .git folder during tests, so they will fail in nix environment
          doCheck = false;

          src = ./.;

          vendorHash = "sha256-qWB4wz4JfxEh4LiixD5JK8/mmF/kmEKwTLh4mqEdDbA=";
          # vendorHash = pkgs.lib.fakeHash;

          meta = with pkgs.lib; {
            description = "CLI management tool";
            homepage = "https://github.com/nxy7/optimus";
            license = licenses.mit;
            maintainers = with maintainers; [ nxyt ];
          };
        };
      });
}
