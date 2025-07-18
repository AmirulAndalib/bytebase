<template>
  <div class="w-full flex flex-col justify-start items-start space-y-4">
    <SQLReviewForResource
      v-if="databaseChangeMode === DatabaseChangeMode.PIPELINE"
      ref="sqlReviewForResourceRef"
      :resource="project.name"
      :allow-edit="allowEdit"
    />
    <RestrictIssueCreationConfigure
      v-if="databaseChangeMode === DatabaseChangeMode.PIPELINE"
      ref="restrictIssueCreationConfigureRef"
      :resource="project.name"
      :allow-edit="allowEdit"
    />
    <AccessControlConfigure
      ref="accessControlConfigureRef"
      :resource="project.name"
      :allow-edit="allowEdit"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import AccessControlConfigure from "@/components/EnvironmentForm/AccessControlConfigure.vue";
import RestrictIssueCreationConfigure from "@/components/GeneralSetting/RestrictIssueCreationConfigure.vue";
import { SQLReviewForResource } from "@/components/SQLReview";
import { useAppFeature } from "@/store";
import type { ComposedProject } from "@/types";
import { DatabaseChangeMode } from "@/types/proto-es/v1/setting_service_pb";

defineProps<{
  project: ComposedProject;
  allowEdit: boolean;
}>();

const databaseChangeMode = useAppFeature("bb.feature.database-change-mode");

const restrictIssueCreationConfigureRef =
  ref<InstanceType<typeof RestrictIssueCreationConfigure>>();
const accessControlConfigureRef =
  ref<InstanceType<typeof AccessControlConfigure>>();
const sqlReviewForResourceRef =
  ref<InstanceType<typeof SQLReviewForResource>>();

const isDirty = computed(
  () =>
    restrictIssueCreationConfigureRef.value?.isDirty ||
    accessControlConfigureRef.value?.isDirty ||
    sqlReviewForResourceRef.value?.isDirty
);

const onUpdate = async () => {
  if (restrictIssueCreationConfigureRef.value?.isDirty) {
    await restrictIssueCreationConfigureRef.value.update();
  }
  if (sqlReviewForResourceRef.value?.isDirty) {
    await sqlReviewForResourceRef.value.update();
  }
  if (accessControlConfigureRef.value?.isDirty) {
    await accessControlConfigureRef.value.update();
  }
};

const resetState = () => {
  sqlReviewForResourceRef.value?.revert();
  accessControlConfigureRef.value?.revert();
  restrictIssueCreationConfigureRef.value?.revert();
};

defineExpose({
  isDirty,
  update: onUpdate,
  revert: resetState,
});
</script>
