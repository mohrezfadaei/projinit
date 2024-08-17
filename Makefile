# Variables
APP_NAME := projinit
VERSION := 1.0.0
MAINTAINER := Mohammad Reza Fadaei <mohrezfadaei@gmail.com>
DESCRIPTION := Initialize git projects using command lines.

BIN_DIR := build/bin
DEB_DIR := build/deb
RPM_DIR := build/rpm
BUILD_DIR := build

ARCHITECTURES := amd64 arm64
PLATFORMS := linux windows darwin

# Mapping between Go architectures and RPM architectures
RPM_ARCH_MAP := $(if $(filter $(GOARCH),amd64),x86_64,$(if $(filter $(GOARCH),arm64),aarch64,$(GOARCH)))

# Build targets
build: $(foreach arch, $(ARCHITECTURES), $(foreach plat, $(PLATFORMS), build-$(plat)-$(arch)))

build-linux: $(foreach arch, $(ARCHITECTURES), build-linux-$(arch))
build-windows: $(foreach arch, $(ARCHITECTURES), build-windows-$(arch))
build-macos: $(foreach arch, $(ARCHITECTURES), build-darwin-$(arch))

build-all: build-linux build-windows build-macos

# Build binaries for specific platforms and architectures
build-%:
	@echo "Building $(APP_NAME) for GOOS=$(word 1,$(subst -, ,$*)) and GOARCH=$(word 2,$(subst -, ,$*))"
	@mkdir -p $(BIN_DIR)
	@if [ "$(word 1,$(subst -, ,$*))" = "windows" ]; then \
		GOOS=$(word 1,$(subst -, ,$*)) GOARCH=$(word 2,$(subst -, ,$*)) go build -o $(BIN_DIR)/$(APP_NAME)-$(word 1,$(subst -, ,$*))-$(word 2,$(subst -, ,$*)).exe main.go; \
	else \
		GOOS=$(word 1,$(subst -, ,$*)) GOARCH=$(word 2,$(subst -, ,$*)) go build -o $(BIN_DIR)/$(APP_NAME)-$(word 1,$(subst -, ,$*))-$(word 2,$(subst -, ,$*)) main.go; \
	fi

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
	mv $(DEB_PACKAGE).deb $(BIN_DIR)/$(APP_NAME)-$*.deb

# Package .rpm for specific architectures
package-rpm: $(foreach arch, $(ARCHITECTURES), package-rpm-$(arch))

package-rpm-%: build-linux-% 
	@echo "Packaging $(APP_NAME) for architecture $(GOARCH)"
	$(eval BIN_PATH := $(BIN_DIR)/$(APP_NAME)-linux-$(GOARCH))
	$(eval RPM_ARCH := $(RPM_ARCH_MAP))
	$(eval RPM_PACKAGE := $(RPM_DIR)/$(RPM_ARCH)/$(APP_NAME))

	@echo "Using RPM_ARCH=$(RPM_ARCH) for architecture $(GOARCH)"

	mkdir -p $(RPM_PACKAGE)/BUILD
	mkdir -p $(RPM_PACKAGE)/RPMS/$(RPM_ARCH)
	mkdir -p $(RPM_PACKAGE)/SOURCES
	mkdir -p $(RPM_PACKAGE)/SPECS
	mkdir -p $(RPM_PACKAGE)/SRPMS

	# Create tarball of the binary inside the expected directory
	mkdir -p $(RPM_PACKAGE)/tmp/$(APP_NAME)-$(VERSION)
	cp $(BIN_PATH) $(RPM_PACKAGE)/tmp/$(APP_NAME)-$(VERSION)/$(APP_NAME)
	tar czvf $(RPM_PACKAGE)/SOURCES/$(APP_NAME)-$(VERSION).tar.gz -C $(RPM_PACKAGE)/tmp $(APP_NAME)-$(VERSION)

	# Create the spec file based on the template
	echo "Name:           $(APP_NAME)" > $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Version:        $(VERSION)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Release:        1%{?dist}" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Summary:        $(DESCRIPTION)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Group:          Development/Tools" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "License:        Apache 2.0" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "URL:            https://github.com/mohrezfadaei/projinit" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "Source0:        %{name}-%{version}.tar.gz" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%description" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "$(DESCRIPTION)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%prep" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%setup -q" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%build" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "# Nothing to build, Go binary is pre-built" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%install" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "install -D -m 0755 $(APP_NAME) %{buildroot}/usr/bin/$(APP_NAME)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "# Prevent `strip` from running" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%global __strip /bin/true" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%files" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "/usr/bin/$(APP_NAME)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "%changelog" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "* Fri Aug 16 2024 $(MAINTAINER) - $(VERSION)" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	echo "- Initial RPM release" >> $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec

	# Build the .rpm package
	rpmbuild --define "_topdir $(abspath $(RPM_PACKAGE))" -ba --target $(RPM_ARCH) $(RPM_PACKAGE)/SPECS/$(APP_NAME).spec
	mv $(RPM_PACKAGE)/RPMS/$(RPM_ARCH)/$(APP_NAME)-$(VERSION)-1.$(RPM_ARCH).rpm $(BIN_DIR)/$(APP_NAME)-$*.rpm

# Package all architectures
package-all: package-deb package-rpm

# Clean build directories
clean:
	rm -rf $(RPM_DIR)
	rm -rf $(DEB_DIR)
	rm -f $(BIN_DIR)/*
	rm -f build/*

# Build all, package, and clean
all: clean build-all package-deb package-rpm
