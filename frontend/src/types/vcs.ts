import { VCSId } from "./id";
import { Principal } from "./principal";

export type VCSType = "GITLAB_SELF_HOST";

export interface VCSConfig {
  type: VCSType;
  name: string;
  instanceUrl: string;
  applicationId: string;
  secret: string;
}

export type VCS = {
  id: VCSId;

  // Standard fields
  creator: Principal;
  createdTs: number;
  updater: Principal;
  updatedTs: number;

  // Domain specific fields
  name: string;
  type: VCSType;
  instanceUrl: string;
  apiUrl: string;
  applicationId: string;
  secret: string;
};

export type VCSCreate = {
  // Domain specific fields
  name: string;
  type: VCSType;
  instanceUrl: string;
  applicationId: string;
  secret: string;
};

export type VCSPatch = {
  // Domain specific fields
  name?: string;
  applicationId?: string;
  secret?: string;
};

export type VCSTokenCreate = {
  code: string;
  redirectUrl: string;
};

export type VCSFileCommit = {
  id: string;
  title: string;
  message: string;
  createdTs: number;
  url: string;
  authorName: string;
  added: string;
};

export type VCSPushEvent = {
  vcsType: VCSType;
  ref: string;
  repositoryId: string;
  repositoryUrl: string;
  repositoryFullPath: string;
  authorName: string;
  fileCommit: VCSFileCommit;
};

export function isValidVCSApplicationIdOrSecret(str: string): boolean {
  return /^[a-zA-Z0-9_]{64}$/.test(str);
}
