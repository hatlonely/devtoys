<script>
  import "beercss";
  import "material-dynamic-colors";
  import { Base64Decode, Base64Encode } from "../wailsjs/go/apps/Base64TextApp";
  import copy from "copy-to-clipboard";

  let decodedText = "";
  let encodedText = "";
  let operation = "encode";
  let type_ = "std";
  let raw = false;
  let decodeError = "";

  let pasted = false;
  let copied = false;

  let operation_infos = [
    {
      text: "编码",
      value: "encode",
    },
    {
      text: "解码",
      value: "decode",
    },
  ];

  let type_infos = [
    {
      text: "URL编码",
      value: "url",
    },
    {
      text: "STD编码",
      value: "std",
    },
  ];

  function calculate() {
    if (operation === "encode") {
      Base64Encode(decodedText, type_, raw).then((result) => {
        encodedText = result;
      });
    } else {
      Base64Decode(encodedText, type_, raw)
        .then((result) => {
          decodedText = result;
          decodeError = "";
        })
        .catch((err) => {
          decodeError = err;
        });
    }
    copied = false;
    pasted = false;
  }

  $: raw, calculate();
  $: type_, calculate();
  $: operation, calculate();
</script>

<main class="responsive">
  <h3>Base64文本编解码</h3>

  <div class="grid border round center abs">
    <div class="s3 field middle-align no-margin">
      <nav class="absolute center">
        {#each operation_infos as info}
          <label class="radio">
            <input
              type="radio"
              name={info.text}
              value={info.value}
              bind:group={operation}
            />
            <span>{info.text}</span>
          </label>
        {/each}
      </nav>
    </div>
    <div class="s3 field middle-align no-margin">
      <nav class="absolute center">
        {#each type_infos as info}
          <label class="radio">
            <input
              type="radio"
              name={info.text}
              value={info.value}
              bind:group={type_}
            />
            <span>{info.text}</span>
          </label>
        {/each}
      </nav>
    </div>

    <div class="s3 field middle-align no-margin">
      <nav class="absolute center">
        <div class="s1">RAW</div>
        <label class="switch">
          <input type="checkbox" bind:checked={raw} />
          <span />
        </label>
      </nav>
    </div>

    <div class="s3 field middle-align no-margin">
      <nav class="absolute center">
        <!-- 粘贴按钮 -->
        <button
          class="border square round small"
          on:click={(e) => {
            navigator.clipboard.readText().then((text) => {
              if (operation === "encode") {
                decodedText = text;
              } else {
                encodedText = text;
              }
              calculate();
              pasted = true;
            });
          }}
          on:focusout={(e) => {
            pasted = false;
          }}
        >
          <i>{pasted ? "done" : "content_paste"}</i>
          <div class="tooltip">粘贴</div>
        </button>
        <!-- 复制按钮 -->
        <button
          class="border square round small"
          on:click={(e) => {
            if (operation === "encode") {
              copy(encodedText);
            } else {
              copy(decodedText);
            }
            copied = true;
          }}
          on:focusout={(e) => {
            pasted = false;
          }}
        >
          <i>{copied ? "done" : "content_copy"}</i>
          <div class="tooltip">复制</div>
        </button>
        <!-- 清空按钮 -->
        <button
          class="border square round small"
          on:click={(e) => {
            if (operation === "encode") {
              decodedText = "";
            } else {
              encodedText = "";
            }
            calculate();
          }}
        >
          <i>{copied ? "done" : "delete"}</i>
          <div class="tooltip">清空</div>
        </button>
      </nav>
    </div>
  </div>

  <div class="medium-space" />

  <div
    class={"field textarea label border" +
      (operation === "decode" ? " fill" : "")}
  >
    <textarea
      bind:value={decodedText}
      name="decodedTextArea"
      disabled={operation === "decode"}
      on:change={calculate}
    />
    <label for="decodedTextArea">编码前字符串</label>
  </div>

  <div
    class={"field textarea label border" +
      (decodeError ? " invalid" : "") +
      (operation === "encode" ? " fill" : "")}
  >
    <textarea
      bind:value={encodedText}
      name="encodedTextArea"
      disabled={operation === "encode"}
      on:change={calculate}
    />
    <label for="encodedTextArea">解码后字符串</label>
    <span class="error">{decodeError}</span>
  </div>
</main>
