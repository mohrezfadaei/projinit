# Variables
APP_NAME := projinit
VERSION := 1.0.0
MAINTAINER := Mohammad Reza Fadaei <mohrezfadaei@gmail.com>
DESCRIPTION := A CLI tool to initialize projects with LICENSE, README.md, and .gitignore files.

BIN_DIR := build/bin
DEB_DIR := build/deb
RPM_DIR := build/rpm
BUILD_DIR := build

# Architectures and platforms
ARCHITECTURES := amd64 arm64
PLATFORMS := linux windows darwin

# Mapping between Go architectures and RPM architectures
RPM_ARCH_MAP := $(if $(filter $*,amd64),x86_64,$(if $(filter $*,arm64),aarch64,$*))

# Build all binaries
build: $(foreach arch, $(ARCHITECTURES), $(foreach plat, $(PLATFORMS), build-$(plat)-$(arch)))

build-linux: $(foreach arch, $(ARCHITECTURES), build-linux-$(arch))
build-windows: $(foreach arch, $(ARCHITECTURES), build-windows-$(arch))
build-macos: $(foreach arch, $(ARCHITECTURES), build-darwin-$(arch))

build-all: build-linux build-windows build-macos

# Build binaries for specific platforms and architectures
build-%:
	@echo "Building $(APP_NAME) for GOOS=$(word 1,$(subst -, ,$*)) and GOARCH=$(word 2,$(subst -, ,$*))"
	GOOS=$(word 1,$(subst -, ,$*)) GOARCH=$(word 2,$(subst -, ,$*)) go build -o $(BIN_DIR)/$(APP_NAME)-$(word 1,$(subst -, ,$*))-$(word 2,$(subst -, ,$*)) main.go

# Package .deb for specific architectures
package-deb: $(foreach arch, $(ARCHITECTURES), package-deb-$(arch))

package-deb-%: build-linux-%
	@echo "Packaging $(APP_NAME) for architecture $*"
	$(eval DEB_PACKAGE := $(DEB_DIR)/$*/$(APP_NAME))
	$(eval BIN_PATH := $(BIN_DIR)/$(APP_NAME)-linux-$*)

	mkdir -p $(DEB_PACKAGE)/DEBIAN
	mkdir -p $(DEB_PACKAGE)/usr/local/bin
	mkdir -p $(DEB_PACKAGE)/usr/share/doc/$(APP_NAME)

	cp $(BIN_PATH) $(DEB_PACKAGE)/usr/local/bin/$(APP_NAME)
	cp README.md $(DEB_PACKAGE)/usr/share/doc/$(APP_NAME)/

	# Control file creation
	echo "Package: $(APP_NAME)" > $(DEB_PACKAGE)/DEBIAN/control
	echo "Version: $(VERSION)" >> $(DEB_PACKAGE)/DEBIAN/control
	echo "Section: base" >> $(DEB_PACKAGE)/DEBIAN/control
	echo "Priority: optional" >> $(DEB_PACKAGE)/DEBIAN/control
	echo "Architecture: $*" >> $(DEB_PACKAGE)/DEBIAN/control
	echo "Maintainer: $(MAINTAINER)" >> $(DEB_PACKAGE)/DEBIAN/control
	echo "Description: $(DESCRIPTION)" >> $(DEB_PACKAGE)/DEBIAN/control

	# Build the .deb package
	dpkg-deb --build $(DEB_PACKAGE)
	mv $(DEB_PACKAGE).deb $(BUILD_DIR)/$(APP_NAME)-$*.deb

# Package .rpm for specific architectures
package-rpm: $(foreach arch, $(ARCHITECTURES), package-rpm-$(arch))

package-rpm-%: build-linux-%
	@echo "Packaging $(APP_NAME) for architecture $*"
	$(eval RPM_PACKAGE := $(RPM_DIR)/$*/$(APP_NAME))
	$(eval BIN_PATH := $(BIN_DIR)/$(APP_NAME)-linux-$*)
	$(eval RPM_ARCH := $(if $(filter $*,amd64),x86_64,$(if $(filter $*,arm64),aarch64,$*)))

	@echo "Using RPM_ARCH=$(RPM_ARCH) for architecture $*"

	mkdir -p $(RPM_PACKAGE)/BUILD
	mkdir -p $(RPM_PACKAGE)/RPMS/$(RPM_ARCH)
	mkdir -p $(RPM_PACKAGE)/SOURCES
	mkdir -p $(RPM_PACKAGE)/SPECS
	mkdir -p $(RPM_PACKAGE)/SRPMS

	# Create tarball of the binary inside the expected directory
	mkdir -p $(RPM_PACKAGE)/tmp/$(APP_NAME)-$(VERSION)
	cp $(BIN_PATH) $(RPM_PACKAGE)/tmp/$(APP_NAME)-$(VERSION)/$(APP_NAME)
	tar czvf $(RPM_PACKAGE)/SOURCES/$(APP_NAME)-$(VERSION).tar.gz -C $(RPM_PACKAGE)/tmp $(APP_NAME)-$(VERSION)

	# Create the spec file
	echo "Name: $(APP_NAME)" > $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Version: $(VERSION)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Release: 1%{?dist}" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Summary: $(DESCRIPTION)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "License: MIT" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "URL: https://example.com/$(APP_NAME)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Source0: $(APP_NAME)-$(VERSION).tar.gz" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "BuildArch: $(RPM_ARCH)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%description" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "$(DESCRIPTION)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%prep" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%setup -q" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%build" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%install" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "install -D -m 0755 $(APP_NAME) %{buildroot}/usr/local/bin/$(APP_NAME)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%files" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "/usr/local/bin/$(APP_NAME)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%changelog" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "* $(shell LC_TIME=C date '+%a %b %d %Y') $(MAINTAINER) - $(VERSION)-1" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "- Initial RPM release" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec

	# Build the .rpm package
	rpmbuild --define "_topdir $(abspath $(RPM_PACKAGE))" -bb $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	mv $(RPM_PACKAGE)/RPMS/$(RPM_ARCH)/$(APP_NAME)-$(VERSION)-1.$(RPM_ARCH).rpm $(BUILD_DIR)/$(APP_NAME)-$*.rpm

# Clean build and package directories
clean:
	rm -rf $(BUILD_DIR)

# Default target
all: clean build-all package-deb package-rpm
