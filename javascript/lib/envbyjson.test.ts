import { describe, expect, test } from "@jest/globals";

import envbyjson from "./envbyjson";

describe("envbyjson module", () => {
  test("load", () => {
    envbyjson.load("lib/testdata/env.json");
    expect(process.env.A).toBe("got A");
    expect(process.env.B).toBe("got B");
    envbyjson.load("lib/testdata/env2.json", "lib/testdata/env3.json");
    expect(process.env.C).toBe("got C");
    expect(process.env.D).toBe("got D");
  });

  test("load wrong", () => {
    const t = () => {
      envbyjson.load("lib/testdata/env_err.json");
    };
    expect(t).toThrow(Error);
  });

  test("loadProp", () => {
    envbyjson.loadProp("lib/testdata/env4.json", "lib/testdata/env5.json", "P");
    expect(process.env.H).toBe("got H");
    expect(process.env.I).toBe("got I");
  });

  test("loadProp wrong", () => {
    const t = () => {
      envbyjson.loadProp("lib/testdata/env4_err.json", "P");
    };
    expect(t).toThrow(Error);
  });
});
