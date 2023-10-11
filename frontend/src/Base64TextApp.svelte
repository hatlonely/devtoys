<script>
  import "beercss";
  import "material-dynamic-colors";
  import { Base64Decode, Base64Encode } from "../wailsjs/go/apps/Base64TextApp";

  let decodedText = "";
  let encodedText = "";
  let operation = "encode";
  let type_ = "std";
  let raw = false;
  let decodeError = "";

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
      text: "URL",
      value: "url",
    },
    {
      text: "STD",
      value: "std",
    },
  ];
</script>

<main class="responsive">
  <h3>Base64文本编解码</h3>

  <div class="field middle-align no-margin">
    <nav>
      <div class="s1"><b>操作</b></div>
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

  <div class="field middle-align no-margin">
    <nav>
      <div class="s1"><b>类型</b></div>
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

  <div class="field middle-align no-margin">
    <nav>
      <div class="s1"><b>RAW</b></div>
      <label class="switch">
        <input type="checkbox" bind:checked={raw} />
        <span />
      </label>
    </nav>
  </div>

  <nav class="no-space">
    <button class="border left-round">
      <i class="small">check</i>
      <span>Button</span>
    </button>
    <button class="border no-round">
      <i class="small">check</i>
      <span>Button</span>
    </button>
    <button class="border right-round fill">
      <i class="small">check</i>
      <span>Button</span>
    </button>
  </nav>

  <div
    class={"field textarea label border" +
      (operation === "decode" ? " fill" : "")}
  >
    <textarea
      bind:value={decodedText}
      name="decodedTextArea"
      disabled={operation === "decode"}
      on:change={(e) => {
        if (operation === "encode") {
          Base64Encode(decodedText, type_, raw).then((result) => {
            encodedText = result;
          });
        }
      }}
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
      name="encodeTextArea"
      disabled={operation === "encode"}
      on:change={(e) => {
        if (operation === "decode") {
          Base64Decode(encodedText, type_, raw)
            .then((result) => {
              decodedText = result;
              decodeError = "";
            })
            .catch((err) => {
              decodeError = err;
            });
        }
      }}
    />
    <label for="encodeTextArea">解码后字符串</label>
    <span class="error">{decodeError}</span>
  </div>
</main>
