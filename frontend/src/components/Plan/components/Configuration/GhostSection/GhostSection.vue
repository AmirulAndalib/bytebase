<template>
  <div class="flex flex-col items-start gap-1">
    <div
      class="w-full flex flex-row items-center gap-2 whitespace-nowrap flex-wrap"
    >
      <GhostSwitch />
      <div class="textlabel flex items-center gap-x-1 whitespace-nowrap">
        <NTooltip>
          <template #trigger>
            {{ $t("task.online-migration.self") }}
          </template>
          <template #default>
            <i18n-t
              tag="p"
              keypath="issue.migration-mode.online.description"
              class="whitespace-pre-line max-w-[20rem]"
            >
              <template #link>
                <!-- TODO: update docs for mariadb -->
                <LearnMoreLink
                  url="https://docs.bytebase.com/change-database/online-schema-migration-for-mysql?source=console"
                  color="light"
                />
              </template>
            </i18n-t>
          </template>
        </NTooltip>
      </div>
      <NButton
        v-if="enabled && allowChange"
        tag="div"
        size="tiny"
        style="--n-padding: 0 5px"
        @click="showFlagsPanel = true"
      >
        <template #icon>
          <WrenchIcon class="w-4 h-4" />
        </template>
        <template #default>
          {{ $t("task.online-migration.configure") }}
        </template>
      </NButton>
    </div>

    <GhostFlagsPanel
      :show="showFlagsPanel"
      @update:show="(show) => (showFlagsPanel = show)"
    />
  </div>
</template>

<script lang="ts" setup>
import { WrenchIcon } from "lucide-vue-next";
import { NButton, NTooltip } from "naive-ui";
import { ref } from "vue";
import LearnMoreLink from "@/components/LearnMoreLink.vue";
import GhostFlagsPanel from "./GhostFlagsPanel.vue";
import GhostSwitch from "./GhostSwitch.vue";
import { useGhostSettingContext } from "./context";

const showFlagsPanel = ref(false);

const { allowChange, enabled } = useGhostSettingContext();
</script>
