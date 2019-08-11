{ deployment, hostname ? "" }:
let
    deploy = import deployment;
    pkgs = deploy.network.pkgs;
    lib = pkgs.lib;

    hosts = lib.filterAttrs (n: _: n != "network") deploy;

    configuration = { lib, ... }: {
        imports = [
          ./options.nix
          hosts."${hostname}"
        ];

        nixpkgs.pkgs = lib.mkDefault deploy.network.pkgs;
    };
in
{
    info = {
        hosts = with lib; mapAttrsToList (name: v: {
          inherit name deployment;
          targetHost = name;
        }) hosts;
    };

    nixos = import "${pkgs.path}/nixos" { inherit configuration; };
}
