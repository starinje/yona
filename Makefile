
install-nix:
	@echo "Checking for Nix installation..."
	@if ! command -v nix &> /dev/null; then \
		echo "Nix is not installed. Installing Nix..."; \
		curl -L https://nixos.org/nix/install | sh; \
		echo "Nix installed successfully. Please restart your terminal or run 'source ~/.nix-profile/etc/profile.d/nix.sh'"; \
	else \
		echo "Nix is already installed."; \
	fi