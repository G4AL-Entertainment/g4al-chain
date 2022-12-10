// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgUpdateDeveloper } from "./types/g4alchain/permission/tx";
import { MsgCreateAdministrator } from "./types/g4alchain/permission/tx";
import { MsgUpdateAdministrator } from "./types/g4alchain/permission/tx";
import { MsgCreateDeveloper } from "./types/g4alchain/permission/tx";


export { MsgUpdateDeveloper, MsgCreateAdministrator, MsgUpdateAdministrator, MsgCreateDeveloper };

type sendMsgUpdateDeveloperParams = {
  value: MsgUpdateDeveloper,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateAdministratorParams = {
  value: MsgCreateAdministrator,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateAdministratorParams = {
  value: MsgUpdateAdministrator,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateDeveloperParams = {
  value: MsgCreateDeveloper,
  fee?: StdFee,
  memo?: string
};


type msgUpdateDeveloperParams = {
  value: MsgUpdateDeveloper,
};

type msgCreateAdministratorParams = {
  value: MsgCreateAdministrator,
};

type msgUpdateAdministratorParams = {
  value: MsgUpdateAdministrator,
};

type msgCreateDeveloperParams = {
  value: MsgCreateDeveloper,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgUpdateDeveloper({ value, fee, memo }: sendMsgUpdateDeveloperParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateDeveloper: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateDeveloper({ value: MsgUpdateDeveloper.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateDeveloper: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateAdministrator({ value, fee, memo }: sendMsgCreateAdministratorParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateAdministrator: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateAdministrator({ value: MsgCreateAdministrator.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateAdministrator: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateAdministrator({ value, fee, memo }: sendMsgUpdateAdministratorParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateAdministrator: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateAdministrator({ value: MsgUpdateAdministrator.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateAdministrator: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateDeveloper({ value, fee, memo }: sendMsgCreateDeveloperParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateDeveloper: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateDeveloper({ value: MsgCreateDeveloper.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateDeveloper: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgUpdateDeveloper({ value }: msgUpdateDeveloperParams): EncodeObject {
			try {
				return { typeUrl: "/g4alentertainment.g4alchain.permission.MsgUpdateDeveloper", value: MsgUpdateDeveloper.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateDeveloper: Could not create message: ' + e.message)
			}
		},
		
		msgCreateAdministrator({ value }: msgCreateAdministratorParams): EncodeObject {
			try {
				return { typeUrl: "/g4alentertainment.g4alchain.permission.MsgCreateAdministrator", value: MsgCreateAdministrator.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateAdministrator: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateAdministrator({ value }: msgUpdateAdministratorParams): EncodeObject {
			try {
				return { typeUrl: "/g4alentertainment.g4alchain.permission.MsgUpdateAdministrator", value: MsgUpdateAdministrator.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateAdministrator: Could not create message: ' + e.message)
			}
		},
		
		msgCreateDeveloper({ value }: msgCreateDeveloperParams): EncodeObject {
			try {
				return { typeUrl: "/g4alentertainment.g4alchain.permission.MsgCreateDeveloper", value: MsgCreateDeveloper.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateDeveloper: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			G4AlentertainmentG4AlchainPermission: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;