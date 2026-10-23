class AttackToolingProcessor extends BaseProcessorV2 {
  inspectRequestV2(metaInfo: MetaInfo, request: Request, tags: Set<string>): InspectResultV2 {
    // Tag requests from the "nikto" web server scanner.
    if (request.header.has("user-agent") && request.header.get("user-agent") === "nikto") {
      tags.add("ATTACK_TOOLING");
    }

    return InspectResultV2.from(tags);
  }

  enforceRequestV2(metaInfo: MetaInfo, request: Request, tags: Set<string>): EnforceResultV2 {
    if (tags.has("ATTACK_TOOLING")) {
      return EnforceResultV2.BLOCK;
    }

    return EnforceResultV2.ALLOW;
  }
}

registerV2(new AttackToolingProcessor());
