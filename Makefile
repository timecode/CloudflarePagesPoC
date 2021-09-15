#!make
SHELL := /bin/sh
THIS_FILE := $(lastword $(MAKEFILE_LIST))

BUILD_DIR ?= "build"
VERSION_FILE ?= "VERSION"
MAJOR_MINOR_PATCH_VERSION ?= "0.0.0"

LOCAL_SERVER ?= "http://local.shadowcryptic.com:1313"

THIS_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
THIS_DIR := $(dir $(THIS_PATH))
BUILD_DIR_DEV = ${BUILD_DIR}/development
BUILD_DIR_PROD = ${BUILD_DIR}/production
DEPLOY_DIR = "public"
MINIFY_SWICTH ?= "--minify"

##############################################################################

# https://github.com/jeffsp/makefile_help
.PHONY: help # Generate list of targets with descriptions
help:
	@ echo "Available commands:" ;\
		grep "^.PHONY: .* #" Makefile | sort | sed "s/.PHONY: \(.*\) # \(.*\)/\1	\2/" | expand -t20

.PHONY: list # [ALIAS] help
list : help

.PHONY: dev # [ALIAS] serve
dev : serve

.PHONY: local # [ALIAS] serve
local : serve

.PHONY: serve # Serve the website locally
serve :
	@ hugo server \
		--gc \
		--navigateToChanged \
		--print-mem \
		--cleanDestinationDir \
		--buildDrafts \
		--minify \
		--baseURL ${LOCAL_SERVER}

.PHONY: test # Run test(s) using locally served site
test :
	@ muffet \
		--exclude=.*fonts.gstatic.com.* \
		--exclude=.*shadowcryptic.com.* \
		--exclude=.*googletagmanager.com.* \
		--exclude=.*music.apple.com.* \
		--exclude=.*groovetech.com.* \
		--follow-sitemap-xml \
		${LOCAL_SERVER}

.PHONY: test-verbose # Run verbose test(s) using locally served site
test-verbose :
	@ muffet \
		--exclude=.*fonts.gstatic.com.* \
		--exclude=.*shadowcryptic.com.* \
		--exclude=.*googletagmanager.com.* \
		--exclude=.*music.apple.com.* \
		--exclude=.*groovetech.com.* \
		--follow-sitemap-xml \
		--verbose \
		${LOCAL_SERVER}

.PHONY: theme-reset # Reset base theme to head
theme-reset :
	@ cd themes/CodeIT; \
		git checkout master && \
    	git fetch --all && \
		git reset --hard origin/master \

.PHONY: clean # Clean all build artifacts
clean : clean-deploy
	@ rm -rf ${BUILD_DIR} && \
		echo "... cleared whole build directory '${BUILD_DIR}'"

clean-deploy :
	@ rm -rf ${DEPLOY_DIR} && \
		echo "... cleared deployment directory '${DEPLOY_DIR}'"

.PHONY: clean-dev # Clean dev build artifacts
clean-dev :
	@ rm -rf ${BUILD_DIR_DEV} && \
		echo "... cleared dev build directory '${BUILD_DIR_DEV}'"

.PHONY: clean-prod # Clean prod build artifacts
clean-prod :
	@ rm -rf ${BUILD_DIR_PROD} && \
		echo "... cleared prod build directory '${BUILD_DIR_PROD}'"

.PHONY: build # [ALIAS] build-dev
build : build-dev

.PHONY: build-dev-nomin # Build dev artifacts without minification
build-dev-nomin :
	@ MINIFY_SWICTH="" $(MAKE) --file=$(THIS_FILE) build-dev

.PHONY: build-dev # Build dev artifacts
build-dev : clean-dev
	@ hugo \
		--gc \
		--environment development \
		--destination ${BUILD_DIR_DEV} \
		${MINIFY_SWICTH} && \
		rm -f ${BUILD_DIR_DEV}/sitemap.xml

.PHONY: build-prod # Build prod artifacts
build-prod : clean-prod
	@ hugo \
		--gc \
		--environment production \
		--destination ${BUILD_DIR_PROD} \
		--MINIFY_SWICTH

.PHONY: deploy-ready-dev # Create a dev Deploy-ready directory
deploy-ready-dev : clean-deploy
	@ if [ ! -d ${THIS_DIR}${BUILD_DIR_DEV}/ ]; then \
			echo "Nothing built yet, so building first..."; \
			$(MAKE) --file=$(THIS_FILE) build-dev; \
		else \
			echo "Using existing build artifacts from ${BUILD_DIR_DEV}"; \
		fi ;\
		rsync -a ${BUILD_DIR_DEV}/ ${DEPLOY_DIR} && \
		echo "... files copied to '${DEPLOY_DIR}'"

.PHONY: deploy-ready-dev # Create a prod Deploy-ready directory
deploy-ready-prod : clean-deploy
	@ if [ ! -d ${THIS_DIR}${BUILD_DIR_PROD}/ ]; then \
			echo "Nothing built yet, so building first..."; \
			$(MAKE) --file=$(THIS_FILE) build-prod; \
		else \
			echo "Using existing build artifacts from ${BUILD_DIR_PROD}"; \
		fi ;\
		rsync -a ${BUILD_DIR_PROD}/ ${DEPLOY_DIR} && \
		echo "... files copied to '${DEPLOY_DIR}'"

.PHONY: deploy-cloudflare-worker # Deploy Cloudflare worker
deploy-cloudflare-worker :
	@ echo "... updating cloudflare worker" && \
		go run ./gocode/cmd/generate-cloudflare-worker && \
		echo "... updated cloudflare worker"

.PHONY: deploy # Deploy
cloudflare-deploy :
	@ hugo && \
		$(MAKE) --file=$(THIS_FILE) deploy-cloudflare-worker
