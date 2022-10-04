import type { NextApiRequest, NextApiResponse } from "next";
import { credentials, ServiceError } from "@grpc/grpc-js";

import { UserServiceClient } from "../../codegen/proto/user_grpc_pb";
import { GetUserRequest, GetUserResponse } from "../../codegen/proto/user_pb";

const Request = new GetUserRequest();
const Client = new UserServiceClient(
  "localhost:50051",
  credentials.createInsecure()
);

export type UserApiResponse =
  | { ok: true; user: GetUserResponse.AsObject["user"] }
  | { ok: false; error: ServiceError };

export default function handler(
  apiReq: NextApiRequest,
  apiRes: NextApiResponse<UserApiResponse>
) {
  const { id } = JSON.parse(apiReq.body);
  Request.setUserid(id);
  Client.getUser(Request, (grpcErr, grpcRes) => {
    if (grpcErr) {
      apiRes.status(500).json({ ok: false, error: grpcErr });
    } else {
      const { user } = grpcRes.toObject();
      apiRes.status(200).json({ ok: true, user });
    }
  });
}
