// @generated by protoc-gen-es v2.5.2
// @generated from file v1/issue_service.proto (package bytebase.v1, syntax proto3)
/* eslint-disable */

import { enumDesc, fileDesc, messageDesc, serviceDesc, tsEnum } from "@bufbuild/protobuf/codegenv2";
import { file_google_api_annotations } from "../google/api/annotations_pb";
import { file_google_api_client } from "../google/api/client_pb";
import { file_google_api_field_behavior } from "../google/api/field_behavior_pb";
import { file_google_api_resource } from "../google/api/resource_pb";
import { file_google_protobuf_duration, file_google_protobuf_field_mask, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_type_expr } from "../google/type/expr_pb";
import { file_v1_annotation } from "./annotation_pb";

/**
 * Describes the file v1/issue_service.proto.
 */
export const file_v1_issue_service = /*@__PURE__*/
  fileDesc("ChZ2MS9pc3N1ZV9zZXJ2aWNlLnByb3RvEgtieXRlYmFzZS52MSJKCg9HZXRJc3N1ZVJlcXVlc3QSKAoEbmFtZRgBIAEoCUIa4EEC+kEUChJieXRlYmFzZS5jb20vSXNzdWUSDQoFZm9yY2UYAiABKAgiagoSQ3JlYXRlSXNzdWVSZXF1ZXN0EiwKBnBhcmVudBgBIAEoCUIc4EEC+kEWChRieXRlYmFzZS5jb20vUHJvamVjdBImCgVpc3N1ZRgCIAEoCzISLmJ5dGViYXNlLnYxLklzc3VlQgPgQQIihwEKEUxpc3RJc3N1ZXNSZXF1ZXN0EiwKBnBhcmVudBgBIAEoCUIc4EEC+kEWChRieXRlYmFzZS5jb20vUHJvamVjdBIRCglwYWdlX3NpemUYAiABKAUSEgoKcGFnZV90b2tlbhgDIAEoCRIOCgZmaWx0ZXIYBCABKAkSDQoFcXVlcnkYBSABKAkiUQoSTGlzdElzc3Vlc1Jlc3BvbnNlEiIKBmlzc3VlcxgBIAMoCzISLmJ5dGViYXNlLnYxLklzc3VlEhcKD25leHRfcGFnZV90b2tlbhgCIAEoCSJwChNTZWFyY2hJc3N1ZXNSZXF1ZXN0EhMKBnBhcmVudBgBIAEoCUID4EECEhEKCXBhZ2Vfc2l6ZRgCIAEoBRISCgpwYWdlX3Rva2VuGAMgASgJEg4KBmZpbHRlchgEIAEoCRINCgVxdWVyeRgFIAEoCSJTChRTZWFyY2hJc3N1ZXNSZXNwb25zZRIiCgZpc3N1ZXMYASADKAsyEi5ieXRlYmFzZS52MS5Jc3N1ZRIXCg9uZXh0X3BhZ2VfdG9rZW4YAiABKAkiiQEKElVwZGF0ZUlzc3VlUmVxdWVzdBI9CgVpc3N1ZRgBIAEoCzISLmJ5dGViYXNlLnYxLklzc3VlQhrgQQL6QRQKEmJ5dGViYXNlLmNvbS9Jc3N1ZRI0Cgt1cGRhdGVfbWFzaxgCIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE1hc2tCA+BBAiKYAQoeQmF0Y2hVcGRhdGVJc3N1ZXNTdGF0dXNSZXF1ZXN0EiwKBnBhcmVudBgBIAEoCUIc4EEC+kEWChRieXRlYmFzZS5jb20vUHJvamVjdBIOCgZpc3N1ZXMYAiADKAkSKAoGc3RhdHVzGAMgASgOMhguYnl0ZWJhc2UudjEuSXNzdWVTdGF0dXMSDgoGcmVhc29uGAQgASgJIiEKH0JhdGNoVXBkYXRlSXNzdWVzU3RhdHVzUmVzcG9uc2UiUAoTQXBwcm92ZUlzc3VlUmVxdWVzdBIoCgRuYW1lGAEgASgJQhrgQQL6QRQKEmJ5dGViYXNlLmNvbS9Jc3N1ZRIPCgdjb21tZW50GAIgASgJIk8KElJlamVjdElzc3VlUmVxdWVzdBIoCgRuYW1lGAEgASgJQhrgQQL6QRQKEmJ5dGViYXNlLmNvbS9Jc3N1ZRIPCgdjb21tZW50GAIgASgJIlAKE1JlcXVlc3RJc3N1ZVJlcXVlc3QSKAoEbmFtZRgBIAEoCUIa4EEC+kEUChJieXRlYmFzZS5jb20vSXNzdWUSDwoHY29tbWVudBgCIAEoCSLqCAoFSXNzdWUSDAoEbmFtZRgBIAEoCRINCgV0aXRsZRgDIAEoCRITCgtkZXNjcmlwdGlvbhgEIAEoCRIlCgR0eXBlGAUgASgOMhcuYnl0ZWJhc2UudjEuSXNzdWUuVHlwZRIoCgZzdGF0dXMYBiABKA4yGC5ieXRlYmFzZS52MS5Jc3N1ZVN0YXR1cxIuCglhcHByb3ZlcnMYCSADKAsyGy5ieXRlYmFzZS52MS5Jc3N1ZS5BcHByb3ZlchI5ChJhcHByb3ZhbF90ZW1wbGF0ZXMYCiADKAsyHS5ieXRlYmFzZS52MS5BcHByb3ZhbFRlbXBsYXRlEh0KFWFwcHJvdmFsX2ZpbmRpbmdfZG9uZRgLIAEoCBIeChZhcHByb3ZhbF9maW5kaW5nX2Vycm9yGAwgASgJEhQKB2NyZWF0b3IYDiABKAlCA+BBAxI0CgtjcmVhdGVfdGltZRgPIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCA+BBAxI0Cgt1cGRhdGVfdGltZRgQIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCA+BBAxIMCgRwbGFuGBEgASgJEg8KB3JvbGxvdXQYEiABKAkSMAoNZ3JhbnRfcmVxdWVzdBgTIAEoCzIZLmJ5dGViYXNlLnYxLkdyYW50UmVxdWVzdBIRCglyZWxlYXNlcnMYFCADKAkSMAoKcmlza19sZXZlbBgVIAEoDjIcLmJ5dGViYXNlLnYxLklzc3VlLlJpc2tMZXZlbBJCChF0YXNrX3N0YXR1c19jb3VudBgWIAMoCzInLmJ5dGViYXNlLnYxLklzc3VlLlRhc2tTdGF0dXNDb3VudEVudHJ5Eg4KBmxhYmVscxgXIAMoCRqcAQoIQXBwcm92ZXISMgoGc3RhdHVzGAEgASgOMiIuYnl0ZWJhc2UudjEuSXNzdWUuQXBwcm92ZXIuU3RhdHVzEhEKCXByaW5jaXBhbBgCIAEoCSJJCgZTdGF0dXMSFgoSU1RBVFVTX1VOU1BFQ0lGSUVEEAASCwoHUEVORElORxABEgwKCEFQUFJPVkVEEAISDAoIUkVKRUNURUQQAxo2ChRUYXNrU3RhdHVzQ291bnRFbnRyeRILCgNrZXkYASABKAkSDQoFdmFsdWUYAiABKAU6AjgBIlkKBFR5cGUSFAoQVFlQRV9VTlNQRUNJRklFRBAAEhMKD0RBVEFCQVNFX0NIQU5HRRABEhEKDUdSQU5UX1JFUVVFU1QQAhITCg9EQVRBQkFTRV9FWFBPUlQQAyJICglSaXNrTGV2ZWwSGgoWUklTS19MRVZFTF9VTlNQRUNJRklFRBAAEgcKA0xPVxABEgwKCE1PREVSQVRFEAISCAoESElHSBADOjrqQTcKEmJ5dGViYXNlLmNvbS9Jc3N1ZRIhcHJvamVjdHMve3Byb2plY3R9L2lzc3Vlcy97aXNzdWV9SgQIAhADSgQIBxAISgQICBAJIn8KDEdyYW50UmVxdWVzdBIMCgRyb2xlGAEgASgJEgwKBHVzZXIYAiABKAkSJAoJY29uZGl0aW9uGAMgASgLMhEuZ29vZ2xlLnR5cGUuRXhwchItCgpleHBpcmF0aW9uGAQgASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uIl8KEEFwcHJvdmFsVGVtcGxhdGUSJwoEZmxvdxgBIAEoCzIZLmJ5dGViYXNlLnYxLkFwcHJvdmFsRmxvdxINCgV0aXRsZRgCIAEoCRITCgtkZXNjcmlwdGlvbhgDIAEoCSI4CgxBcHByb3ZhbEZsb3cSKAoFc3RlcHMYASADKAsyGS5ieXRlYmFzZS52MS5BcHByb3ZhbFN0ZXAilgEKDEFwcHJvdmFsU3RlcBIsCgR0eXBlGAEgASgOMh4uYnl0ZWJhc2UudjEuQXBwcm92YWxTdGVwLlR5cGUSKAoFbm9kZXMYAiADKAsyGS5ieXRlYmFzZS52MS5BcHByb3ZhbE5vZGUiLgoEVHlwZRIUChBUWVBFX1VOU1BFQ0lGSUVEEAASBwoDQUxMEAESBwoDQU5ZEAIiegoMQXBwcm92YWxOb2RlEiwKBHR5cGUYASABKA4yHi5ieXRlYmFzZS52MS5BcHByb3ZhbE5vZGUuVHlwZRIMCgRyb2xlGAIgASgJIi4KBFR5cGUSFAoQVFlQRV9VTlNQRUNJRklFRBAAEhAKDEFOWV9JTl9HUk9VUBABIm0KGExpc3RJc3N1ZUNvbW1lbnRzUmVxdWVzdBIqCgZwYXJlbnQYASABKAlCGuBBAvpBFAoSYnl0ZWJhc2UuY29tL0lzc3VlEhEKCXBhZ2Vfc2l6ZRgCIAEoBRISCgpwYWdlX3Rva2VuGAMgASgJImcKGUxpc3RJc3N1ZUNvbW1lbnRzUmVzcG9uc2USMQoOaXNzdWVfY29tbWVudHMYASADKAsyGS5ieXRlYmFzZS52MS5Jc3N1ZUNvbW1lbnQSFwoPbmV4dF9wYWdlX3Rva2VuGAIgASgJInkKGUNyZWF0ZUlzc3VlQ29tbWVudFJlcXVlc3QSKgoGcGFyZW50GAEgASgJQhrgQQL6QRQKEmJ5dGViYXNlLmNvbS9Jc3N1ZRIwCg1pc3N1ZV9jb21tZW50GAIgASgLMhkuYnl0ZWJhc2UudjEuSXNzdWVDb21tZW50Iq8BChlVcGRhdGVJc3N1ZUNvbW1lbnRSZXF1ZXN0EioKBnBhcmVudBgBIAEoCUIa4EEC+kEUChJieXRlYmFzZS5jb20vSXNzdWUSMAoNaXNzdWVfY29tbWVudBgCIAEoCzIZLmJ5dGViYXNlLnYxLklzc3VlQ29tbWVudBI0Cgt1cGRhdGVfbWFzaxgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE1hc2tCA+BBAiKuDAoMSXNzdWVDb21tZW50EgwKBG5hbWUYASABKAkSDwoHY29tbWVudBgCIAEoCRIPCgdwYXlsb2FkGAMgASgJEjQKC2NyZWF0ZV90aW1lGAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEID4EEDEjQKC3VwZGF0ZV90aW1lGAUgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEID4EEDEhQKB2NyZWF0b3IYByABKAlCA+BBAxI2CghhcHByb3ZhbBgIIAEoCzIiLmJ5dGViYXNlLnYxLklzc3VlQ29tbWVudC5BcHByb3ZhbEgAEj0KDGlzc3VlX3VwZGF0ZRgJIAEoCzIlLmJ5dGViYXNlLnYxLklzc3VlQ29tbWVudC5Jc3N1ZVVwZGF0ZUgAEjcKCXN0YWdlX2VuZBgKIAEoCzIiLmJ5dGViYXNlLnYxLklzc3VlQ29tbWVudC5TdGFnZUVuZEgAEjsKC3Rhc2tfdXBkYXRlGAsgASgLMiQuYnl0ZWJhc2UudjEuSXNzdWVDb21tZW50LlRhc2tVcGRhdGVIABJGChF0YXNrX3ByaW9yX2JhY2t1cBgMIAEoCzIpLmJ5dGViYXNlLnYxLklzc3VlQ29tbWVudC5UYXNrUHJpb3JCYWNrdXBIABqQAQoIQXBwcm92YWwSOQoGc3RhdHVzGAEgASgOMikuYnl0ZWJhc2UudjEuSXNzdWVDb21tZW50LkFwcHJvdmFsLlN0YXR1cyJJCgZTdGF0dXMSFgoSU1RBVFVTX1VOU1BFQ0lGSUVEEAASCwoHUEVORElORxABEgwKCEFQUFJPVkVEEAISDAoIUkVKRUNURUQQAxr1AgoLSXNzdWVVcGRhdGUSFwoKZnJvbV90aXRsZRgBIAEoCUgAiAEBEhUKCHRvX3RpdGxlGAIgASgJSAGIAQESHQoQZnJvbV9kZXNjcmlwdGlvbhgDIAEoCUgCiAEBEhsKDnRvX2Rlc2NyaXB0aW9uGAQgASgJSAOIAQESMgoLZnJvbV9zdGF0dXMYBSABKA4yGC5ieXRlYmFzZS52MS5Jc3N1ZVN0YXR1c0gEiAEBEjAKCXRvX3N0YXR1cxgGIAEoDjIYLmJ5dGViYXNlLnYxLklzc3VlU3RhdHVzSAWIAQESEwoLZnJvbV9sYWJlbHMYCSADKAkSEQoJdG9fbGFiZWxzGAogAygJQg0KC19mcm9tX3RpdGxlQgsKCV90b190aXRsZUITChFfZnJvbV9kZXNjcmlwdGlvbkIRCg9fdG9fZGVzY3JpcHRpb25CDgoMX2Zyb21fc3RhdHVzQgwKCl90b19zdGF0dXNKBAgHEAhKBAgIEAkaGQoIU3RhZ2VFbmQSDQoFc3RhZ2UYASABKAkapwIKClRhc2tVcGRhdGUSDQoFdGFza3MYASADKAkSFwoKZnJvbV9zaGVldBgCIAEoCUgAiAEBEhUKCHRvX3NoZWV0GAMgASgJSAGIAQESQwoJdG9fc3RhdHVzGAYgASgOMisuYnl0ZWJhc2UudjEuSXNzdWVDb21tZW50LlRhc2tVcGRhdGUuU3RhdHVzSAKIAQEiawoGU3RhdHVzEhYKElNUQVRVU19VTlNQRUNJRklFRBAAEgsKB1BFTkRJTkcQARILCgdSVU5OSU5HEAISCAoERE9ORRADEgoKBkZBSUxFRBAEEgsKB1NLSVBQRUQQBRIMCghDQU5DRUxFRBAGQg0KC19mcm9tX3NoZWV0QgsKCV90b19zaGVldEIMCgpfdG9fc3RhdHVzGtcBCg9UYXNrUHJpb3JCYWNrdXASDAoEdGFzaxgBIAEoCRI/CgZ0YWJsZXMYAiADKAsyLy5ieXRlYmFzZS52MS5Jc3N1ZUNvbW1lbnQuVGFza1ByaW9yQmFja3VwLlRhYmxlEhoKDW9yaWdpbmFsX2xpbmUYAyABKAVIAIgBARIQCghkYXRhYmFzZRgEIAEoCRINCgVlcnJvchgFIAEoCRomCgVUYWJsZRIOCgZzY2hlbWEYASABKAkSDQoFdGFibGUYAiABKAlCEAoOX29yaWdpbmFsX2xpbmVCBwoFZXZlbnRKBAgGEAcqTQoLSXNzdWVTdGF0dXMSHAoYSVNTVUVfU1RBVFVTX1VOU1BFQ0lGSUVEEAASCAoET1BFThABEggKBERPTkUQAhIMCghDQU5DRUxFRBADMtgPCgxJc3N1ZVNlcnZpY2USgAEKCEdldElzc3VlEhwuYnl0ZWJhc2UudjEuR2V0SXNzdWVSZXF1ZXN0GhIuYnl0ZWJhc2UudjEuSXNzdWUiQtpBBG5hbWWK6jANYmIuaXNzdWVzLmdldJDqMAGC0+STAiASHi92MS97bmFtZT1wcm9qZWN0cy8qL2lzc3Vlcy8qfRKcAQoLQ3JlYXRlSXNzdWUSHy5ieXRlYmFzZS52MS5DcmVhdGVJc3N1ZVJlcXVlc3QaEi5ieXRlYmFzZS52MS5Jc3N1ZSJY2kEMcGFyZW50LGlzc3VliuowEGJiLmlzc3Vlcy5jcmVhdGWQ6jABmOowAYLT5JMCJzoFaXNzdWUiHi92MS97cGFyZW50PXByb2plY3RzLyp9L2lzc3VlcxKUAQoKTGlzdElzc3VlcxIeLmJ5dGViYXNlLnYxLkxpc3RJc3N1ZXNSZXF1ZXN0Gh8uYnl0ZWJhc2UudjEuTGlzdElzc3Vlc1Jlc3BvbnNlIkXaQQZwYXJlbnSK6jAOYmIuaXNzdWVzLmxpc3SQ6jABgtPkkwIgEh4vdjEve3BhcmVudD1wcm9qZWN0cy8qfS9pc3N1ZXMSmgEKDFNlYXJjaElzc3VlcxIgLmJ5dGViYXNlLnYxLlNlYXJjaElzc3Vlc1JlcXVlc3QaIS5ieXRlYmFzZS52MS5TZWFyY2hJc3N1ZXNSZXNwb25zZSJFiuowDWJiLmlzc3Vlcy5nZXSQ6jACgtPkkwIqOgEqIiUvdjEve3BhcmVudD1wcm9qZWN0cy8qfS9pc3N1ZXM6c2VhcmNoEqcBCgtVcGRhdGVJc3N1ZRIfLmJ5dGViYXNlLnYxLlVwZGF0ZUlzc3VlUmVxdWVzdBoSLmJ5dGViYXNlLnYxLklzc3VlImPaQRFpc3N1ZSx1cGRhdGVfbWFza4rqMBBiYi5pc3N1ZXMudXBkYXRlkOowAZjqMAGC0+STAi06BWlzc3VlMiQvdjEve2lzc3VlLm5hbWU9cHJvamVjdHMvKi9pc3N1ZXMvKn0SwAEKEUxpc3RJc3N1ZUNvbW1lbnRzEiUuYnl0ZWJhc2UudjEuTGlzdElzc3VlQ29tbWVudHNSZXF1ZXN0GiYuYnl0ZWJhc2UudjEuTGlzdElzc3VlQ29tbWVudHNSZXNwb25zZSJc2kEGcGFyZW50iuowFWJiLmlzc3VlQ29tbWVudHMubGlzdJDqMAGC0+STAjASLi92MS97cGFyZW50PXByb2plY3RzLyovaXNzdWVzLyp9L2lzc3VlQ29tbWVudHMS0gEKEkNyZWF0ZUlzc3VlQ29tbWVudBImLmJ5dGViYXNlLnYxLkNyZWF0ZUlzc3VlQ29tbWVudFJlcXVlc3QaGS5ieXRlYmFzZS52MS5Jc3N1ZUNvbW1lbnQiedpBFHBhcmVudCxpc3N1ZV9jb21tZW50iuowF2JiLmlzc3VlQ29tbWVudHMuY3JlYXRlkOowAZjqMAGC0+STAjk6DWlzc3VlX2NvbW1lbnQiKC92MS97cGFyZW50PXByb2plY3RzLyovaXNzdWVzLyp9OmNvbW1lbnQS3wEKElVwZGF0ZUlzc3VlQ29tbWVudBImLmJ5dGViYXNlLnYxLlVwZGF0ZUlzc3VlQ29tbWVudFJlcXVlc3QaGS5ieXRlYmFzZS52MS5Jc3N1ZUNvbW1lbnQihQHaQSBwYXJlbnQsaXNzdWVfY29tbWVudCx1cGRhdGVfbWFza4rqMBdiYi5pc3N1ZUNvbW1lbnRzLnVwZGF0ZZDqMAGY6jABgtPkkwI5Og1pc3N1ZV9jb21tZW50MigvdjEve3BhcmVudD1wcm9qZWN0cy8qL2lzc3Vlcy8qfTpjb21tZW50Es0BChdCYXRjaFVwZGF0ZUlzc3Vlc1N0YXR1cxIrLmJ5dGViYXNlLnYxLkJhdGNoVXBkYXRlSXNzdWVzU3RhdHVzUmVxdWVzdBosLmJ5dGViYXNlLnYxLkJhdGNoVXBkYXRlSXNzdWVzU3RhdHVzUmVzcG9uc2UiV4rqMBBiYi5pc3N1ZXMudXBkYXRlkOowAZjqMAGC0+STAjU6ASoiMC92MS97cGFyZW50PXByb2plY3RzLyp9L2lzc3VlczpiYXRjaFVwZGF0ZVN0YXR1cxJ/CgxBcHByb3ZlSXNzdWUSIC5ieXRlYmFzZS52MS5BcHByb3ZlSXNzdWVSZXF1ZXN0GhIuYnl0ZWJhc2UudjEuSXNzdWUiOZDqMAKY6jABgtPkkwIrOgEqIiYvdjEve25hbWU9cHJvamVjdHMvKi9pc3N1ZXMvKn06YXBwcm92ZRJ8CgtSZWplY3RJc3N1ZRIfLmJ5dGViYXNlLnYxLlJlamVjdElzc3VlUmVxdWVzdBoSLmJ5dGViYXNlLnYxLklzc3VlIjiQ6jACmOowAYLT5JMCKjoBKiIlL3YxL3tuYW1lPXByb2plY3RzLyovaXNzdWVzLyp9OnJlamVjdBJ/CgxSZXF1ZXN0SXNzdWUSIC5ieXRlYmFzZS52MS5SZXF1ZXN0SXNzdWVSZXF1ZXN0GhIuYnl0ZWJhc2UudjEuSXNzdWUiOZDqMAKY6jABgtPkkwIrOgEqIiYvdjEve25hbWU9cHJvamVjdHMvKi9pc3N1ZXMvKn06cmVxdWVzdEI2WjRnaXRodWIuY29tL2J5dGViYXNlL2J5dGViYXNlL2JhY2tlbmQvZ2VuZXJhdGVkLWdvL3YxYgZwcm90bzM", [file_google_api_annotations, file_google_api_client, file_google_api_field_behavior, file_google_api_resource, file_google_protobuf_duration, file_google_protobuf_field_mask, file_google_protobuf_timestamp, file_google_type_expr, file_v1_annotation]);

/**
 * Describes the message bytebase.v1.GetIssueRequest.
 * Use `create(GetIssueRequestSchema)` to create a new message.
 */
export const GetIssueRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 0);

/**
 * Describes the message bytebase.v1.CreateIssueRequest.
 * Use `create(CreateIssueRequestSchema)` to create a new message.
 */
export const CreateIssueRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 1);

/**
 * Describes the message bytebase.v1.ListIssuesRequest.
 * Use `create(ListIssuesRequestSchema)` to create a new message.
 */
export const ListIssuesRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 2);

/**
 * Describes the message bytebase.v1.ListIssuesResponse.
 * Use `create(ListIssuesResponseSchema)` to create a new message.
 */
export const ListIssuesResponseSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 3);

/**
 * Describes the message bytebase.v1.SearchIssuesRequest.
 * Use `create(SearchIssuesRequestSchema)` to create a new message.
 */
export const SearchIssuesRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 4);

/**
 * Describes the message bytebase.v1.SearchIssuesResponse.
 * Use `create(SearchIssuesResponseSchema)` to create a new message.
 */
export const SearchIssuesResponseSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 5);

/**
 * Describes the message bytebase.v1.UpdateIssueRequest.
 * Use `create(UpdateIssueRequestSchema)` to create a new message.
 */
export const UpdateIssueRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 6);

/**
 * Describes the message bytebase.v1.BatchUpdateIssuesStatusRequest.
 * Use `create(BatchUpdateIssuesStatusRequestSchema)` to create a new message.
 */
export const BatchUpdateIssuesStatusRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 7);

/**
 * Describes the message bytebase.v1.BatchUpdateIssuesStatusResponse.
 * Use `create(BatchUpdateIssuesStatusResponseSchema)` to create a new message.
 */
export const BatchUpdateIssuesStatusResponseSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 8);

/**
 * Describes the message bytebase.v1.ApproveIssueRequest.
 * Use `create(ApproveIssueRequestSchema)` to create a new message.
 */
export const ApproveIssueRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 9);

/**
 * Describes the message bytebase.v1.RejectIssueRequest.
 * Use `create(RejectIssueRequestSchema)` to create a new message.
 */
export const RejectIssueRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 10);

/**
 * Describes the message bytebase.v1.RequestIssueRequest.
 * Use `create(RequestIssueRequestSchema)` to create a new message.
 */
export const RequestIssueRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 11);

/**
 * Describes the message bytebase.v1.Issue.
 * Use `create(IssueSchema)` to create a new message.
 */
export const IssueSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 12);

/**
 * Describes the message bytebase.v1.Issue.Approver.
 * Use `create(Issue_ApproverSchema)` to create a new message.
 */
export const Issue_ApproverSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 12, 0);

/**
 * Describes the enum bytebase.v1.Issue.Approver.Status.
 */
export const Issue_Approver_StatusSchema = /*@__PURE__*/
  enumDesc(file_v1_issue_service, 12, 0, 0);

/**
 * @generated from enum bytebase.v1.Issue.Approver.Status
 */
export const Issue_Approver_Status = /*@__PURE__*/
  tsEnum(Issue_Approver_StatusSchema);

/**
 * Describes the enum bytebase.v1.Issue.Type.
 */
export const Issue_TypeSchema = /*@__PURE__*/
  enumDesc(file_v1_issue_service, 12, 0);

/**
 * @generated from enum bytebase.v1.Issue.Type
 */
export const Issue_Type = /*@__PURE__*/
  tsEnum(Issue_TypeSchema);

/**
 * Describes the enum bytebase.v1.Issue.RiskLevel.
 */
export const Issue_RiskLevelSchema = /*@__PURE__*/
  enumDesc(file_v1_issue_service, 12, 1);

/**
 * @generated from enum bytebase.v1.Issue.RiskLevel
 */
export const Issue_RiskLevel = /*@__PURE__*/
  tsEnum(Issue_RiskLevelSchema);

/**
 * Describes the message bytebase.v1.GrantRequest.
 * Use `create(GrantRequestSchema)` to create a new message.
 */
export const GrantRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 13);

/**
 * Describes the message bytebase.v1.ApprovalTemplate.
 * Use `create(ApprovalTemplateSchema)` to create a new message.
 */
export const ApprovalTemplateSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 14);

/**
 * Describes the message bytebase.v1.ApprovalFlow.
 * Use `create(ApprovalFlowSchema)` to create a new message.
 */
export const ApprovalFlowSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 15);

/**
 * Describes the message bytebase.v1.ApprovalStep.
 * Use `create(ApprovalStepSchema)` to create a new message.
 */
export const ApprovalStepSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 16);

/**
 * Describes the enum bytebase.v1.ApprovalStep.Type.
 */
export const ApprovalStep_TypeSchema = /*@__PURE__*/
  enumDesc(file_v1_issue_service, 16, 0);

/**
 * Type of the ApprovalStep
 * ALL means every node must be approved to proceed.
 * ANY means approving any node will proceed.
 *
 * @generated from enum bytebase.v1.ApprovalStep.Type
 */
export const ApprovalStep_Type = /*@__PURE__*/
  tsEnum(ApprovalStep_TypeSchema);

/**
 * Describes the message bytebase.v1.ApprovalNode.
 * Use `create(ApprovalNodeSchema)` to create a new message.
 */
export const ApprovalNodeSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 17);

/**
 * Describes the enum bytebase.v1.ApprovalNode.Type.
 */
export const ApprovalNode_TypeSchema = /*@__PURE__*/
  enumDesc(file_v1_issue_service, 17, 0);

/**
 * Type of the ApprovalNode.
 * type determines who should approve this node.
 * ANY_IN_GROUP means the ApprovalNode can be approved by an user from our predefined user group.
 * See GroupValue below for the predefined user groups.
 *
 * @generated from enum bytebase.v1.ApprovalNode.Type
 */
export const ApprovalNode_Type = /*@__PURE__*/
  tsEnum(ApprovalNode_TypeSchema);

/**
 * Describes the message bytebase.v1.ListIssueCommentsRequest.
 * Use `create(ListIssueCommentsRequestSchema)` to create a new message.
 */
export const ListIssueCommentsRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 18);

/**
 * Describes the message bytebase.v1.ListIssueCommentsResponse.
 * Use `create(ListIssueCommentsResponseSchema)` to create a new message.
 */
export const ListIssueCommentsResponseSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 19);

/**
 * Describes the message bytebase.v1.CreateIssueCommentRequest.
 * Use `create(CreateIssueCommentRequestSchema)` to create a new message.
 */
export const CreateIssueCommentRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 20);

/**
 * Describes the message bytebase.v1.UpdateIssueCommentRequest.
 * Use `create(UpdateIssueCommentRequestSchema)` to create a new message.
 */
export const UpdateIssueCommentRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 21);

/**
 * Describes the message bytebase.v1.IssueComment.
 * Use `create(IssueCommentSchema)` to create a new message.
 */
export const IssueCommentSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 22);

/**
 * Describes the message bytebase.v1.IssueComment.Approval.
 * Use `create(IssueComment_ApprovalSchema)` to create a new message.
 */
export const IssueComment_ApprovalSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 22, 0);

/**
 * Describes the enum bytebase.v1.IssueComment.Approval.Status.
 */
export const IssueComment_Approval_StatusSchema = /*@__PURE__*/
  enumDesc(file_v1_issue_service, 22, 0, 0);

/**
 * @generated from enum bytebase.v1.IssueComment.Approval.Status
 */
export const IssueComment_Approval_Status = /*@__PURE__*/
  tsEnum(IssueComment_Approval_StatusSchema);

/**
 * Describes the message bytebase.v1.IssueComment.IssueUpdate.
 * Use `create(IssueComment_IssueUpdateSchema)` to create a new message.
 */
export const IssueComment_IssueUpdateSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 22, 1);

/**
 * Describes the message bytebase.v1.IssueComment.StageEnd.
 * Use `create(IssueComment_StageEndSchema)` to create a new message.
 */
export const IssueComment_StageEndSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 22, 2);

/**
 * Describes the message bytebase.v1.IssueComment.TaskUpdate.
 * Use `create(IssueComment_TaskUpdateSchema)` to create a new message.
 */
export const IssueComment_TaskUpdateSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 22, 3);

/**
 * Describes the enum bytebase.v1.IssueComment.TaskUpdate.Status.
 */
export const IssueComment_TaskUpdate_StatusSchema = /*@__PURE__*/
  enumDesc(file_v1_issue_service, 22, 3, 0);

/**
 * @generated from enum bytebase.v1.IssueComment.TaskUpdate.Status
 */
export const IssueComment_TaskUpdate_Status = /*@__PURE__*/
  tsEnum(IssueComment_TaskUpdate_StatusSchema);

/**
 * Describes the message bytebase.v1.IssueComment.TaskPriorBackup.
 * Use `create(IssueComment_TaskPriorBackupSchema)` to create a new message.
 */
export const IssueComment_TaskPriorBackupSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 22, 4);

/**
 * Describes the message bytebase.v1.IssueComment.TaskPriorBackup.Table.
 * Use `create(IssueComment_TaskPriorBackup_TableSchema)` to create a new message.
 */
export const IssueComment_TaskPriorBackup_TableSchema = /*@__PURE__*/
  messageDesc(file_v1_issue_service, 22, 4, 0);

/**
 * Describes the enum bytebase.v1.IssueStatus.
 */
export const IssueStatusSchema = /*@__PURE__*/
  enumDesc(file_v1_issue_service, 0);

/**
 * @generated from enum bytebase.v1.IssueStatus
 */
export const IssueStatus = /*@__PURE__*/
  tsEnum(IssueStatusSchema);

/**
 * @generated from service bytebase.v1.IssueService
 */
export const IssueService = /*@__PURE__*/
  serviceDesc(file_v1_issue_service, 0);

