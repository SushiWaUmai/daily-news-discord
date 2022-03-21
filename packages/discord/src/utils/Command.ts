import { Message } from "discord.js";

export class Command {
  name: string;
  description: string;
  aliases?: string[];
  usage?: string;
  examples?: string[];

  execute: (msg: Message, args: string[]) => void | Promise<void>;

  constructor({
    name,
    description,
    aliases,
    usage,
    examples,
    execute,
  }: CommandSettings) {
    this.name = name;
    this.description = description;
    this.aliases = aliases;
    this.usage = usage;
    this.examples = examples;
    this.execute = execute;
  }
}

export interface CommandSettings {
  name: string;
  description: string;
  aliases?: string[];
  usage?: string;
  tags?: string[];
  examples?: string[];
  execute: (msg: Message, args: string[]) => void | Promise<void>;
}
