<script>
  import "beercss";
  import "material-dynamic-colors";
  import { Base64Decode, Base64Encode } from "../wailsjs/go/apps/Base64TextApp";

  let plainText = "";
  let encodedText = "";
  let operation = "encode";

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

  let decodeError = "";
</script>

<main>
  <h3>Base64文本编解码</h3>

  <div class="field middle-align">
    <nav>
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
      bind:value={plainText}
      name="plainTextArea"
      disabled={operation === "decode"}
      on:change={(e) => {
        if (operation === "encode") {
          Base64Encode(plainText).then((result) => {
            encodedText = result;
          });
        }
      }}
    />
    <label for="plainTextArea">编码前字符串</label>
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
          Base64Decode(encodedText)
            .then((result) => {
              plainText = result;
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
