###############
### TestNet ###
###############

.PHONY: testnet_supplier_list
testnet_supplier_list: ## List all the staked supplier on TestNet
	poktrolld q supplier list-suppliers --node=$(TESTNET_RPC)

.PHONY: testnet_gateway_list
testnet_gateway_list: ## List all the staked gateways on TestNet
	poktrolld q gateway list-gateway --node=$(TESTNET_RPC)

.PHONY: testnet_app_list
testnet_app_list: ## List all the staked applications on TestNet
	poktrolld q application list-application --node=$(TESTNET_RPC)

.PHONY: testnet_consensus_params
testnet_consensus_params: ## Output consensus parameters
	poktrolld q consensus params --node=$(TESTNET_RPC)

.PHONY: testnet_gov_params
testnet_gov_params: ## Output gov parameters
	poktrolld q gov params --node=$(TESTNET_RPC)

.PHONY: testnet_status
testnet_status: ## Output status of the RPC node (most likely a validator)
	poktrolld status --node=$(TESTNET_RPC) | jq

.PHONY: testnet_height
testnet_height: ## Height of the network from the RPC node point of view
	poktrolld status --node=$(TESTNET_RPC) | jq ".sync_info.latest_block_height"
