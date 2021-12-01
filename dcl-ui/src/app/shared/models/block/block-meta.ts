/**
 * Copyright 2020 DSR Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { JsonObject, JsonProperty } from 'json2typescript';
import { BlockId } from './block-id';
import { BlockHeader } from './block-header';

@JsonObject('BlockMeta')
export class BlockMeta {
  @JsonProperty('block_id', BlockId)
  blockId: BlockId = undefined;

  @JsonProperty('header', BlockHeader)
  header: BlockHeader = undefined;
}
