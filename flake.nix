{
  description = "vorpal";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];

      perSystem = {
        config,
        pkgs,
        ...
      }: let
        name = "example";
        vendorHash = null;
      in {
        devShells = {
          default = pkgs.mkShell {
            inputsFrom = [config.packages.default];
          };
        };

        packages = {
          default = pkgs.buildGoModule {
            inherit name vendorHash;
            src = ./.;
            subPackages = ["cmd/example"];
          };
        };
      };
    };
}
